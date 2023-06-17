package controller

import (
	"net/http"

	"hk4e/node/api"
	"hk4e/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (c *Controller) serverStopState(context *gin.Context) {
	stopServerInfo, err := c.discoveryClient.GetStopServerInfo(context.Request.Context(), &api.GetStopServerInfoReq{ClientIpAddr: ""})
	if err != nil {
		logger.Error("get stop server info error: %v", err)
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	context.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: stopServerInfo})
}

type ServerStopChangeReq struct {
	StopServer bool   `json:"stop_server"`
	StartTime  uint32 `json:"start_time"`
	EndTime    uint32 `json:"end_time"`
}

func (c *Controller) serverStopChange(context *gin.Context) {
	req := new(ServerStopChangeReq)
	err := context.ShouldBindJSON(req)
	if err != nil {
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetStopServerInfo(context.Request.Context(), &api.StopServerInfo{
		StopServer: req.StopServer,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})
	if err != nil {
		logger.Error("set stop server info error: %v", err)
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	context.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}

func (c *Controller) serverWhiteList(context *gin.Context) {
	whiteList, err := c.discoveryClient.GetWhiteList(context.Request.Context(), &api.NullMsg{})
	if err != nil {
		logger.Error("get white list error: %v", err)
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	context.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: whiteList.IpAddrList})
}

type ServerWhiteAdd struct {
	IpAddr string `json:"ip_addr"`
}

func (c *Controller) serverWhiteAdd(context *gin.Context) {
	req := new(ServerWhiteAdd)
	err := context.ShouldBindJSON(req)
	if err != nil {
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetWhiteList(context.Request.Context(), &api.SetWhiteListReq{
		IsAdd:  true,
		IpAddr: req.IpAddr,
	})
	if err != nil {
		logger.Error("set white list error: %v", err)
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	context.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}

type ServerWhiteDel struct {
	IpAddr string `json:"ip_addr"`
}

func (c *Controller) serverWhiteDel(context *gin.Context) {
	req := new(ServerWhiteDel)
	err := context.ShouldBindJSON(req)
	if err != nil {
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetWhiteList(context.Request.Context(), &api.SetWhiteListReq{
		IsAdd:  false,
		IpAddr: req.IpAddr,
	})
	if err != nil {
		logger.Error("set white list error: %v", err)
		context.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	context.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}
