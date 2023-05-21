package main

import (
	"encoding/base64"
	"encoding/hex"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	hk4egatenet "hk4e/gate/net"
	"hk4e/pkg/endec"
	"hk4e/pkg/object"
	"hk4e/pkg/random"
	"hk4e/robot/net"

	pb "google.golang.org/protobuf/proto"

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

	dispatchInfo, err := login.GetDispatchInfo(config.GetConfig().Hk4eRobot.RegionListUrl,
		config.GetConfig().Hk4eRobot.RegionListParam,
		config.GetConfig().Hk4eRobot.CurRegionUrl,
		config.GetConfig().Hk4eRobot.CurRegionParam,
		config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("get dispatch info error: %v", err)
		time.Sleep(time.Second)
		return
	}

	if config.GetConfig().Hk4eRobot.DosEnable {
		dosBatchNum := int(config.GetConfig().Hk4eRobot.DosBatchNum)
		for i := 0; i < int(config.GetConfig().Hk4eRobot.DosTotalNum); i += dosBatchNum {
			wg := new(sync.WaitGroup)
			wg.Add(dosBatchNum)
			for j := 0; j < dosBatchNum; j++ {
				go httpLogin(config.GetConfig().Hk4eRobot.Account+"_"+strconv.Itoa(i+j), dispatchInfo, wg)
			}
			wg.Wait()
		}
	} else {
		httpLogin(config.GetConfig().Hk4eRobot.Account, dispatchInfo, nil)
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

func httpLogin(account string, dispatchInfo *login.DispatchInfo, wg *sync.WaitGroup) {
	defer func() {
		if config.GetConfig().Hk4eRobot.DosEnable {
			wg.Done()
		}
	}()
	accountInfo, err := login.AccountLogin(config.GetConfig().Hk4eRobot.LoginSdkUrl, account, config.GetConfig().Hk4eRobot.Password)
	if err != nil {
		logger.Error("account login error: %v", err)
		return
	}
	logger.Info("robot http login ok, account: %v", account)
	go func() {
		for {
			gateLogin(account, dispatchInfo, accountInfo)
			if !config.GetConfig().Hk4eRobot.DosLoopLogin {
				break
			}
			time.Sleep(time.Second)
			continue
		}
	}()
}

func gateLogin(account string, dispatchInfo *login.DispatchInfo, accountInfo *login.AccountInfo) {
	session, err := login.GateLogin(dispatchInfo, accountInfo, config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("gate login error: %v", err)
		return
	}
	logger.Info("robot gate login ok, account: %v", account)
	clientVersionHashData, err := hex.DecodeString(
		endec.Sha1Str(config.GetConfig().Hk4eRobot.ClientVersion + session.ClientVersionRandomKey + "mhy2020"),
	)
	if err != nil {
		logger.Error("gen clientVersionHashData error: %v", err)
		return
	}
	checksumClientVersion := strings.Split(config.GetConfig().Hk4eRobot.ClientVersion, "_")[0]
	session.SendMsg(cmd.PlayerLoginReq, &proto.PlayerLoginReq{
		AccountType:           1,
		SubChannelId:          1,
		LanguageType:          2,
		PlatformType:          3,
		Checksum:              "$008094416f86a051270e64eb0b405a38825",
		ChecksumClientVersion: checksumClientVersion,
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
	tickCounter := uint64(0)
	pingSeq := uint32(0)
	enterSceneDone := false
	sceneBeginTime := uint32(0)
	bornPos := new(proto.Vector)
	currPos := new(proto.Vector)
	avatarEntityId := uint32(0)
	moveRot := random.GetRandomFloat32(0.0, 359.9)
	moveReliableSeq := uint32(0)
	for {
		select {
		case <-ticker.C:
			tickCounter++
			if config.GetConfig().Hk4eRobot.ClientMoveEnable {
				if enterSceneDone {
					for {
						dx := float32(float64(config.GetConfig().Hk4eRobot.ClientMoveSpeed) * math.Cos(float64(moveRot/360.0*2*math.Pi)))
						dz := float32(float64(config.GetConfig().Hk4eRobot.ClientMoveSpeed) * math.Sin(float64(moveRot/360.0*2*math.Pi)))
						if currPos.X-dx > bornPos.X+float32(config.GetConfig().Hk4eRobot.ClientMoveRangeExt) ||
							currPos.Z-dz > bornPos.Z+float32(config.GetConfig().Hk4eRobot.ClientMoveRangeExt) ||
							currPos.X-dx < bornPos.X-float32(config.GetConfig().Hk4eRobot.ClientMoveRangeExt) ||
							currPos.Z-dz < bornPos.Z-float32(config.GetConfig().Hk4eRobot.ClientMoveRangeExt) {
							moveRot = random.GetRandomFloat32(0.0, 359.9)
							continue
						}
						currPos.X -= dx
						currPos.Z -= dz
						break
					}
					moveReliableSeq += 100
					entityMoveInfo := &proto.EntityMoveInfo{
						EntityId: avatarEntityId,
						MotionInfo: &proto.MotionInfo{
							Pos:    currPos,
							Rot:    &proto.Vector{X: 0.0, Y: moveRot, Z: 0.0},
							Speed:  new(proto.Vector),
							State:  proto.MotionState_MOTION_RUN,
							RefPos: new(proto.Vector),
						},
						SceneTime:   uint32(time.Now().UnixMilli()) - sceneBeginTime,
						ReliableSeq: moveReliableSeq,
						IsReliable:  true,
					}
					logger.Debug("EntityMoveInfo: %v, account: %v", entityMoveInfo, account)
					combatData, err := pb.Marshal(entityMoveInfo)
					if err != nil {
						logger.Error("marshal EntityMoveInfo error: %v, account: %v", err, account)
						continue
					}
					combatInvocationsNotify := &proto.CombatInvocationsNotify{
						InvokeList: []*proto.CombatInvokeEntry{{
							CombatData:   combatData,
							ForwardType:  proto.ForwardType_FORWARD_TO_ALL_EXCEPT_CUR,
							ArgumentType: proto.CombatTypeArgument_ENTITY_MOVE,
						}},
					}
					var combatInvocationsNotifyPb pb.Message = combatInvocationsNotify
					if config.GetConfig().Hk4e.ClientProtoProxyEnable {
						clientProtoObj := hk4egatenet.GetClientProtoObjByName("CombatInvocationsNotify", session.ClientCmdProtoMap)
						if clientProtoObj == nil {
							continue
						}
						_, err := object.CopyProtoBufSameField(clientProtoObj, combatInvocationsNotify)
						if err != nil {
							continue
						}
						hk4egatenet.ConvServerPbDataToClient(clientProtoObj, session.ClientCmdProtoMap)
						combatInvocationsNotifyPb = clientProtoObj
					}
					body, err := pb.Marshal(combatInvocationsNotifyPb)
					if err != nil {
						logger.Error("marshal CombatInvocationsNotify error: %v, account: %v", err, account)
						continue
					}
					unionCmdNotify := &proto.UnionCmdNotify{
						CmdList: []*proto.UnionCmd{{
							Body:      body,
							MessageId: cmd.CombatInvocationsNotify,
						}},
					}
					if config.GetConfig().Hk4e.ClientProtoProxyEnable {
						unionCmdNotify.CmdList[0].MessageId = uint32(session.ClientCmdProtoMap.GetClientCmdIdByCmdName("CombatInvocationsNotify"))
					}
					session.SendMsg(cmd.UnionCmdNotify, unionCmdNotify)
				}
			}
			if tickCounter%5 != 0 {
				continue
			}
			pingSeq++
			// 通过这个接口发消息给服务器
			session.SendMsg(cmd.PingReq, &proto.PingReq{
				ClientTime: uint32(time.Now().Unix()),
				Seq:        pingSeq,
			})
		case protoMsg := <-session.RecvChan:
			// 从这个管道接收服务器发来的消息
			switch protoMsg.CmdId {
			case cmd.PlayerLoginRsp:
				rsp := protoMsg.PayloadMessage.(*proto.PlayerLoginRsp)
				if rsp.Retcode != 0 {
					logger.Error("login fail, retCode: %v, account: %v", rsp.Retcode, account)
					return
				}
				logger.Info("robot gs login ok, account: %v", account)
			case cmd.DoSetPlayerBornDataNotify:
				session.SendMsg(cmd.SetPlayerBornDataReq, &proto.SetPlayerBornDataReq{
					AvatarId: 10000007,
					NickName: account,
				})
			case cmd.PlayerEnterSceneNotify:
				ntf := protoMsg.PayloadMessage.(*proto.PlayerEnterSceneNotify)
				bornPos.X, bornPos.Y, bornPos.Z = ntf.Pos.X, ntf.Pos.Y, ntf.Pos.Z
				currPos.X, currPos.Y, currPos.Z = ntf.Pos.X, ntf.Pos.Y, ntf.Pos.Z
				session.SendMsg(cmd.EnterSceneReadyReq, &proto.EnterSceneReadyReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.EnterSceneReadyRsp:
				ntf := protoMsg.PayloadMessage.(*proto.EnterSceneReadyRsp)
				session.SendMsg(cmd.SceneInitFinishReq, &proto.SceneInitFinishReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.SceneInitFinishRsp:
				ntf := protoMsg.PayloadMessage.(*proto.SceneInitFinishRsp)
				session.SendMsg(cmd.EnterSceneDoneReq, &proto.EnterSceneDoneReq{EnterSceneToken: ntf.EnterSceneToken})
			case cmd.EnterSceneDoneRsp:
				ntf := protoMsg.PayloadMessage.(*proto.EnterSceneDoneRsp)
				enterSceneDone = true
				sceneBeginTime = uint32(time.Now().UnixMilli())
				session.SendMsg(cmd.PostEnterSceneReq, &proto.PostEnterSceneReq{EnterSceneToken: ntf.EnterSceneToken})
				if config.GetConfig().Hk4eRobot.DosLoopLogin {
					return
				}
			case cmd.SceneEntityAppearNotify:
				ntf := protoMsg.PayloadMessage.(*proto.SceneEntityAppearNotify)
				for _, sceneEntityInfo := range ntf.EntityList {
					if sceneEntityInfo.EntityType != proto.ProtEntityType_PROT_ENTITY_AVATAR {
						continue
					}
					avatarEntityId = sceneEntityInfo.EntityId
				}
			}
		case <-session.DeadEvent:
			logger.Info("robot exit, account: %v", account)
			return
		}
	}
}
