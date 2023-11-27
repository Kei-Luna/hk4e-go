package game

import (
	"time"

	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/object"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

/************************************************** 接口请求 **************************************************/

func (g *Game) GetAllMailReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetAllMailReq)

	rsp := &proto.GetAllMailRsp{
		MailList:    g.PacketMailList(player.PlayerId),
		IsCollected: req.IsCollected,
		IsTruncated: false,
	}
	g.SendMsg(cmd.GetAllMailRsp, player.PlayerId, player.ClientSeq, rsp)
}

func (g *Game) GetAllMailNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.GetAllMailNotify)

	pbMailList := make([]*proto.MailData, 0)
	if player.ClientVersion >= 400 {
		pbMailList = g.PacketMailList(player.PlayerId)
	}

	g.SendMsg(cmd.GetAllMailResultNotify, player.PlayerId, player.ClientSeq, &proto.GetAllMailResultNotify{
		Transaction:    g.NewTransaction(player.PlayerId),
		MailList:       pbMailList,
		PageIndex:      1,
		TotalPageCount: 1,
		IsCollected:    ntf.IsCollected,
	})
}

func (g *Game) DelMailReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DelMailReq)

	g.DelPlayerMail(player.PlayerId, req.MailIdList)

	rsp := &proto.DelMailRsp{
		MailIdList: req.MailIdList,
	}
	g.SendMsg(cmd.DelMailRsp, player.PlayerId, player.ClientSeq, rsp)
}

func (g *Game) GetMailItemReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetMailItemReq)

	rsp := &proto.GetMailItemRsp{
		MailIdList: req.MailIdList,
		ItemList:   nil,
	}
	g.SendMsg(cmd.GetMailItemRsp, player.PlayerId, player.ClientSeq, rsp)
}

func (g *Game) ReadMailNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.ReadMailNotify)
	g.ReadPlayerMail(player.PlayerId, ntf.MailIdList)
}

func (g *Game) ChangeMailStarNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.ChangeMailStarNotify)
	g.StarPlayerMail(player.PlayerId, ntf.MailIdList, ntf.IsStar)
}

/************************************************** 游戏功能 **************************************************/

var MailIdSeq uint32 = 0

func (g *Game) AddPlayerMail(userId uint32, title string, content string) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	MailIdSeq++
	mail := &model.Mail{
		MailId:     MailIdSeq,
		Title:      title,
		Content:    content,
		Sender:     "flswld",
		SendTime:   uint32(time.Now().Unix()),
		ExpireTime: uint32(time.Now().Add(time.Hour * 24).Unix()),
		IsRead:     false,
		IsStar:     false,
	}
	player.MailMap[mail.MailId] = mail
	ntf := &proto.MailChangeNotify{
		MailList:      []*proto.MailData{g.PacketMail(mail)},
		DelMailIdList: nil,
	}
	g.SendMsg(cmd.MailChangeNotify, player.PlayerId, player.ClientSeq, ntf)
}

func (g *Game) DelPlayerMail(userId uint32, mailIdList []uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	for _, mailId := range mailIdList {
		delete(player.MailMap, mailId)
	}
	ntf := &proto.MailChangeNotify{
		MailList:      nil,
		DelMailIdList: mailIdList,
	}
	g.SendMsg(cmd.MailChangeNotify, player.PlayerId, player.ClientSeq, ntf)
}

func (g *Game) ReadPlayerMail(userId uint32, mailIdList []uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	for _, mailId := range mailIdList {
		mail, exist := player.MailMap[mailId]
		if !exist {
			return
		}
		mail.IsRead = true
	}
}

func (g *Game) StarPlayerMail(userId uint32, mailIdList []uint32, isStar bool) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	pbMailList := make([]*proto.MailData, 0)
	for _, mailId := range mailIdList {
		mail, exist := player.MailMap[mailId]
		if !exist {
			return
		}
		mail.IsStar = isStar
		pbMail := g.PacketMail(mail)
		if pbMail == nil {
			return
		}
		pbMailList = append(pbMailList, pbMail)
	}
	ntf := &proto.MailChangeNotify{
		MailList:      pbMailList,
		DelMailIdList: nil,
	}
	g.SendMsg(cmd.MailChangeNotify, player.PlayerId, player.ClientSeq, ntf)
}

/************************************************** 打包封装 **************************************************/

func (g *Game) PacketMail(mail *model.Mail) *proto.MailData {
	return &proto.MailData{
		MailId: mail.MailId,
		MailTextContent: &proto.MailTextContent{
			Title:   mail.Title,
			Content: mail.Content,
			Sender:  mail.Sender,
		},
		ItemList:        nil,
		SendTime:        mail.SendTime,
		ExpireTime:      mail.ExpireTime,
		Importance:      uint32(object.ConvBoolToInt64(mail.IsStar)),
		IsRead:          mail.IsRead,
		IsAttachmentGot: false,
		ConfigId:        0,
		ArgumentList:    nil,
		CollectState:    proto.MailCollectState_MAIL_NOT_COLLECTIBLE,
	}
}

func (g *Game) PacketMailList(userId uint32) []*proto.MailData {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return nil
	}
	pbMailList := make([]*proto.MailData, 0)
	for _, mail := range player.MailMap {
		pbMail := g.PacketMail(mail)
		if pbMail == nil {
			return nil
		}
		pbMailList = append(pbMailList, pbMail)
	}
	return pbMailList
}
