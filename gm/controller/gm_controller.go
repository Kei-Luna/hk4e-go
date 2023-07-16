package controller

import (
	"net/http"

	"hk4e/common/rpc"
	"hk4e/gs/api"
	"hk4e/pkg/logger"

	"github.com/gin-gonic/gin"
)

type GmCmdReq struct {
	FuncName  string   `json:"func_name"`
	ParamList []string `json:"param_list"`
	GsId      uint32   `json:"gs_id"`
}

func (c *Controller) gmCmd(ctx *gin.Context) {
	gmCmdReq := new(GmCmdReq)
	err := ctx.ShouldBindJSON(gmCmdReq)
	if err != nil {
		logger.Error("parse json error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	logger.Info("GmCmdReq: %v", gmCmdReq)
	c.gmClientMapLock.RLock()
	gmClient, exist := c.gmClientMap[gmCmdReq.GsId]
	c.gmClientMapLock.RUnlock()
	if !exist {
		var err error = nil
		gmClient, err = rpc.NewGMClient(gmCmdReq.GsId)
		if err != nil {
			logger.Error("new gm client error: %v", err)
			ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
			return
		}
		c.gmClientMapLock.Lock()
		c.gmClientMap[gmCmdReq.GsId] = gmClient
		c.gmClientMapLock.Unlock()
	}
	rep, err := gmClient.Cmd(ctx.Request.Context(), &api.CmdRequest{
		FuncName:  gmCmdReq.FuncName,
		ParamList: gmCmdReq.ParamList,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, rep)
}
