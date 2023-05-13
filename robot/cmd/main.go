package main

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"hk4e/pkg/endec"
	"hk4e/robot/net"

	"hk4e/common/config"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
	"hk4e/robot/login"
)

func main() {
	config.InitConfig("application.toml")
	logger.InitLogger("robot")
	defer func() {
		logger.CloseLogger()
	}()

	// // DPDK模式需开启
	// err := engine.InitEngine("00:0C:29:3E:3E:DF", "192.168.199.199", "255.255.255.0", "192.168.199.1")
	// if err != nil {
	// 	panic(err)
	// }
	// engine.RunEngine([]int{0, 1, 2, 3}, 4, 1, "0.0.0.0")
	// time.Sleep(time.Second * 30)

	if config.GetConfig().Hk4eRobot.DosEnable {
		for i := 0; i < int(config.GetConfig().Hk4eRobot.DosNum); i++ {
			go httpLogin(config.GetConfig().Hk4eRobot.Account + "_" + strconv.Itoa(i))
		}
	} else {
		httpLogin(config.GetConfig().Hk4eRobot.Account)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// // DPDK模式需开启
			// engine.StopEngine()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func httpLogin(account string) {
	logger.Info("robot start, account: %v", account)
	dispatchInfo, err := login.GetDispatchInfo(config.GetConfig().Hk4eRobot.RegionListUrl,
		config.GetConfig().Hk4eRobot.RegionListParam,
		config.GetConfig().Hk4eRobot.CurRegionUrl,
		config.GetConfig().Hk4eRobot.CurRegionParam,
		config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("get dispatch info error: %v", err)
		return
	}
	accountInfo, err := login.AccountLogin(config.GetConfig().Hk4eRobot.LoginSdkUrl, account, config.GetConfig().Hk4eRobot.Password)
	if err != nil {
		logger.Error("account login error: %v", err)
		return
	}
	for {
		gateLogin(account, dispatchInfo, accountInfo)
		if !config.GetConfig().Hk4eRobot.DosLoopLogin {
			break
		}
		time.Sleep(time.Second)
		continue
	}
}

func gateLogin(account string, dispatchInfo *login.DispatchInfo, accountInfo *login.AccountInfo) {
	session, err := login.GateLogin(dispatchInfo, accountInfo, config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("gate login error: %v", err)
		return
	}
	clientVersionHashData, err := hex.DecodeString(
		endec.Sha1Str(config.GetConfig().Hk4eRobot.ClientVersion + session.ClientVersionRandomKey + "mhy2020"),
	)
	if err != nil {
		logger.Error("gen clientVersionHashData error: %v", err)
		return
	}
	session.SendMsg(cmd.PlayerLoginReq, &proto.PlayerLoginReq{
		AccountType:           1,
		SubChannelId:          1,
		LanguageType:          2,
		PlatformType:          3,
		Checksum:              "$008094416f86a051270e64eb0b405a38825",
		ChecksumClientVersion: "CNRELWin3.2.0",
		ClientDataVersion:     11793813,
		ClientVerisonHash:     base64.StdEncoding.EncodeToString(clientVersionHashData),
		ClientVersion:         config.GetConfig().Hk4eRobot.ClientVersion,
		SecurityCmdReply:      session.SecurityCmdBuffer,
		SecurityLibraryMd5:    "574a507ffee2eb6f997d11f71c8ae1fa",
		Token:                 accountInfo.ComboToken,
	})
	clientLogic(account, session)
}

func clientLogic(account string, session *net.Session) {
	ticker := time.NewTicker(time.Second)
	pingSeq := uint32(0)
	for {
		select {
		case <-ticker.C:
			pingSeq++
			// 通过这个接口发消息给服务器
			session.SendMsg(cmd.PingReq, &proto.PingReq{
				ClientTime: uint32(time.Now().Unix()),
				Seq:        pingSeq,
			})
		case protoMsg := <-session.RecvChan:
			// 从这个管道接收服务器发来的消息
			logger.Debug("recv protoMsg: %v", protoMsg)
			switch protoMsg.CmdId {
			case cmd.PlayerLoginRsp:
				rsp := protoMsg.PayloadMessage.(*proto.PlayerLoginRsp)
				logger.Info("login ok, rsp: %v", rsp)
			case cmd.DoSetPlayerBornDataNotify:
				session.SendMsg(cmd.SetPlayerBornDataReq, &proto.SetPlayerBornDataReq{
					AvatarId: 10000007,
					NickName: account,
				})
			case cmd.PlayerEnterSceneNotify:
				ntf := protoMsg.PayloadMessage.(*proto.PlayerEnterSceneNotify)
				session.SendMsg(cmd.EnterSceneReadyReq, &proto.EnterSceneReadyReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.EnterSceneReadyRsp:
				ntf := protoMsg.PayloadMessage.(*proto.EnterSceneReadyRsp)
				session.SendMsg(cmd.SceneInitFinishReq, &proto.SceneInitFinishReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.SceneInitFinishRsp:
				ntf := protoMsg.PayloadMessage.(*proto.SceneInitFinishRsp)
				session.SendMsg(cmd.EnterSceneDoneReq, &proto.EnterSceneDoneReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.EnterSceneDoneRsp:
				ntf := protoMsg.PayloadMessage.(*proto.EnterSceneDoneRsp)
				session.SendMsg(cmd.PostEnterSceneReq, &proto.PostEnterSceneReq{EnterSceneToken: ntf.EnterSceneToken})
				if config.GetConfig().Hk4eRobot.DosLoopLogin {
					return
				}
			}
		case <-session.DeadEvent:
			logger.Info("robot exit, account: %v", account)
			return
		}
	}
}
