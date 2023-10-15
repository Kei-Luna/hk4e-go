package controller

import (
	"net/http"
	"strconv"
	"sync"

	"hk4e/common/config"
	"hk4e/common/rpc"
	"hk4e/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	gmClientMap     map[uint32]*rpc.GMClient
	gmClientMapLock sync.RWMutex
	discoveryClient *rpc.DiscoveryClient
}

func NewController(discoveryClient *rpc.DiscoveryClient) (r *Controller) {
	r = new(Controller)
	r.gmClientMap = make(map[uint32]*rpc.GMClient)
	r.discoveryClient = discoveryClient
	go r.registerRouter()
	return r
}

func (c *Controller) authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		if true {
			// 验证通过
			context.Next()
			return
		}
		// 验证不通过
		context.Abort()
		context.JSON(http.StatusOK, gin.H{
			"code": "10001",
			"msg":  "没有访问权限",
		})
	}
}

type CommonRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (c *Controller) registerRouter() {
	if config.GetConfig().Logger.Level == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	engine.Use(c.authorize())
	engine.POST("/gm/cmd", c.gmCmd)
	engine.GET("/server/stop/state", c.serverStopState)
	engine.POST("/server/stop/change", c.serverStopChange)
	engine.GET("/server/white/list", c.serverWhiteList)
	engine.POST("/server/white/add", c.serverWhiteAdd)
	engine.POST("/server/white/del", c.serverWhiteDel)
	engine.POST("/server/dispatch/cancel", c.serverDispatchCancel)
	port := config.GetConfig().HttpPort
	addr := ":" + strconv.Itoa(int(port))
	err := engine.Run(addr)
	if err != nil {
		logger.Error("gin run error: %v", err)
	}
}
