package client

import (
	"math"
	"time"

	"hk4e/common/config"
	hk4egatenet "hk4e/gate/net"
	"hk4e/pkg/logger"
	"hk4e/pkg/object"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
	"hk4e/robot/net"

	pb "google.golang.org/protobuf/proto"
)

func Logic(account string, session *net.Session) {
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
			case cmd.PlayerDataNotify:
				ntf := protoMsg.PayloadMessage.(*proto.PlayerDataNotify)
				logger.Info("player name: %v", ntf.NickName)
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
					session.Close()
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
			close(session.SendChan)
			return
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
						err := object.CopyProtoBufSameField(clientProtoObj, combatInvocationsNotify)
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
		}
	}
}
