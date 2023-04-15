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

func (c *Controller) gmCmd(context *gin.Context) {
	gmCmdReq := new(GmCmdReq)
	err := context.ShouldBindJSON(gmCmdReq)
	if err != nil {
		logger.Error("parse json error: %v", err)
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
			return
		}
		c.gmClientMapLock.Lock()
		c.gmClientMap[gmCmdReq.GsId] = gmClient
		c.gmClientMapLock.Unlock()
	}
	rep, err := gmClient.Cmd(context.Request.Context(), &api.CmdRequest{
		FuncName:  gmCmdReq.FuncName,
		ParamList: gmCmdReq.ParamList,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, rep)
}
