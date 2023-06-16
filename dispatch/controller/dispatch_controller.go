package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/common/region"
	httpapi "hk4e/dispatch/api"
	"hk4e/node/api"
	"hk4e/pkg/endec"
	"hk4e/pkg/httpclient"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/proto"

	"github.com/gin-gonic/gin"
	pb "google.golang.org/protobuf/proto"
)

func (c *Controller) querySecurityFile(context *gin.Context) {
	// 很早以前2.6.0版本的时候抓包为了完美还原写的 不清楚有没有副作用暂时不要了
	return
	file, err := os.ReadFile("static/security_file")
	if err != nil {
		logger.Error("open security_file error")
		return
	}
	context.Header("Content-type", "text/html; charset=UTF-8")
	_, _ = context.Writer.WriteString(string(file))
}

func (c *Controller) queryRegionList(context *gin.Context) {
	context.Header("Content-type", "text/html; charset=UTF-8")
	var regionListBase64 = ""
	if !config.GetConfig().Hk4e.ForwardModeEnable {
		regionList := region.GetRegionList(c.ec2b)
		regionListData, err := pb.Marshal(regionList)
		if err != nil {
			logger.Error("pb marshal QueryRegionListHttpRsp error: %v", err)
			_, _ = context.Writer.WriteString("500")
			return
		}
		regionListBase64 = base64.StdEncoding.EncodeToString(regionListData)
	} else {
		var err error = nil
		url := context.Request.RequestURI
		param := url[strings.Index(url, "?"):]
		regionListBase64Raw, err := httpclient.GetRaw(config.GetConfig().Hk4e.ForwardRegionUrl + param)
		if err != nil {
			_, _ = context.Writer.WriteString("500")
			return
		}
		regionListDataRaw, err := base64.StdEncoding.DecodeString(regionListBase64Raw)
		if err != nil {
			_, _ = context.Writer.WriteString("500")
			return
		}
		queryRegionListHttpRsp := new(proto.QueryRegionListHttpRsp)
		err = pb.Unmarshal(regionListDataRaw, queryRegionListHttpRsp)
		if err != nil {
			_, _ = context.Writer.WriteString("500")
			return
		}
		logger.Debug("QueryRegionListHttpRsp: %+v", queryRegionListHttpRsp)
		for _, regionSimpleInfo := range queryRegionListHttpRsp.RegionList {
			regionSimpleInfo.DispatchUrl = config.GetConfig().Hk4e.DispatchUrl
		}
		regionListData, err := pb.Marshal(queryRegionListHttpRsp)
		if err != nil {
			_, _ = context.Writer.WriteString("500")
			return
		}
		regionListBase64 = base64.StdEncoding.EncodeToString(regionListData)
	}
	_, _ = context.Writer.WriteString(regionListBase64)
}

func (c *Controller) getClientVersionByName(versionName string) (int, string) {
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		logger.Error("compile regexp error: %v", err)
		return 0, ""
	}
	versionSlice := reg.FindAllString(versionName, -1)
	version := 0
	for index, value := range versionSlice {
		v, err := strconv.Atoi(value)
		if err != nil {
			logger.Error("parse client version error: %v", err)
			return 0, ""
		}
		if v >= 10 {
			// 测试版本
			if index != 2 {
				logger.Error("invalid client version")
				return 0, ""
			}
			v /= 10
		}
		for i := 0; i < 2-index; i++ {
			v *= 10
		}
		version += v
	}
	return version, strconv.Itoa(version)
}

func (c *Controller) queryCurRegion(context *gin.Context) {
	rspError := func() {
		rspContentError := "CAESGE5vdCBGb3VuZCB2ZXJzaW9uIGNvbmZpZw=="
		rspSignError := ""
		rsp := &httpapi.QueryCurRegionRspJson{
			Content: rspContentError,
			Sign:    rspSignError,
		}
		context.JSON(http.StatusOK, rsp)
	}
	versionName := context.Query("version")
	if versionName == "" {
		rspError()
		return
	}
	keyId := context.Query("key_id")
	encPubPrivKey, exist := c.encRsaKeyMap[keyId]
	if !exist {
		logger.Error("can not found key id: %v", keyId)
		rspError()
		return
	}
	version, versionStr := c.getClientVersionByName(versionName)
	if version == 0 {
		rspError()
		return
	}
	addr, err := c.discovery.GetGateServerAddr(context.Request.Context(), &api.GetGateServerAddrReq{
		Version: versionStr,
	})
	if err != nil {
		logger.Error("get gate server addr error: %v", err)
		rspError()
		return
	}
	var regionCurr *proto.QueryCurrRegionHttpRsp = nil
	if !config.GetConfig().Hk4e.ForwardModeEnable {
		regionCurr = region.GetRegionCurr(addr.KcpAddr, int32(addr.KcpPort), c.ec2b)
	} else {
		url := context.Request.RequestURI
		param := url[strings.Index(url, "?"):]
		regionCurrJson, err := httpclient.GetRaw(config.GetConfig().Hk4e.ForwardDispatchUrl + param)
		if err != nil {
			rspError()
			return
		}
		queryCurRegionRspJson := new(httpapi.QueryCurRegionRspJson)
		err = json.Unmarshal([]byte(regionCurrJson), queryCurRegionRspJson)
		if err != nil {
			rspError()
			return
		}
		encryptedRegionInfo, err := base64.StdEncoding.DecodeString(queryCurRegionRspJson.Content)
		if err != nil {
			rspError()
			return
		}
		chunkSize := 256
		regionInfoLength := len(encryptedRegionInfo)
		numChunks := int(math.Ceil(float64(regionInfoLength) / float64(chunkSize)))
		regionCurrData := make([]byte, 0)
		for i := 0; i < numChunks; i++ {
			from := i * chunkSize
			to := int(math.Min(float64((i+1)*chunkSize), float64(regionInfoLength)))
			chunk := encryptedRegionInfo[from:to]
			privKey, err := endec.RsaParsePrivKey(encPubPrivKey)
			if err != nil {
				logger.Error("parse rsa priv key error: %v", err)
				rspError()
				return
			}
			decrypt, err := endec.RsaDecrypt(chunk, privKey)
			if err != nil {
				logger.Error("rsa dec error: %v", err)
				rspError()
				return
			}
			regionCurrData = append(regionCurrData, decrypt...)
		}
		queryCurrRegionHttpRsp := new(proto.QueryCurrRegionHttpRsp)
		err = pb.Unmarshal(regionCurrData, queryCurrRegionHttpRsp)
		if err != nil {
			rspError()
			return
		}
		logger.Debug("QueryCurrRegionHttpRsp: %+v", queryCurrRegionHttpRsp)
		robotServerAppId, err := c.discovery.GetServerAppId(context, &api.GetServerAppIdReq{
			ServerType: api.ROBOT,
		})
		if err != nil {
			logger.Error("get robot server appid error: %v", err)
			rspError()
			return
		}
		ec2b, err := random.LoadEc2bKey(queryCurrRegionHttpRsp.ClientSecretKey)
		if err != nil {
			logger.Error("parse ec2b error: %v", err)
			rspError()
			return
		}
		c.messageQueue.SendToRobot(robotServerAppId.AppId, &mq.NetMsg{
			MsgType: mq.MsgTypeServer,
			EventId: mq.ServerForwardDispatchInfoNotify,
			ServerMsg: &mq.ServerMsg{
				ForwardDispatchInfo: &mq.ForwardDispatchInfo{
					GateIp:      queryCurrRegionHttpRsp.RegionInfo.GateserverIp,
					GatePort:    queryCurrRegionHttpRsp.RegionInfo.GateserverPort,
					DispatchKey: ec2b.XorKey(),
				},
			},
		})
		regionCurr = queryCurrRegionHttpRsp
		regionCurr.ClientSecretKey = c.ec2b.Bytes()
		regionCurr.RegionInfo.GateserverIp = addr.KcpAddr
		regionCurr.RegionInfo.GateserverPort = addr.KcpPort
		endec.Xor(regionCurr.RegionCustomConfigEncrypted, ec2b.XorKey())
		logger.Info("RegionCustomConfigEncrypted: %v", string(regionCurr.RegionCustomConfigEncrypted))
		endec.Xor(regionCurr.RegionCustomConfigEncrypted, c.ec2b.XorKey())
		endec.Xor(regionCurr.ClientRegionCustomConfigEncrypted, ec2b.XorKey())
		logger.Info("ClientRegionCustomConfigEncrypted: %v", string(regionCurr.ClientRegionCustomConfigEncrypted))
		endec.Xor(regionCurr.ClientRegionCustomConfigEncrypted, c.ec2b.XorKey())
	}
	regionCurrData, err := pb.Marshal(regionCurr)
	if err != nil {
		logger.Error("pb marshal QueryCurrRegionHttpRsp error: %v", err)
		rspError()
		return
	}
	if version < 275 {
		context.Header("Content-type", "text/html; charset=UTF-8")
		regionCurrBase64 := base64.StdEncoding.EncodeToString(regionCurrData)
		_, _ = context.Writer.WriteString(regionCurrBase64)
		return
	}
	chunkSize := 256 - 11
	regionInfoLength := len(regionCurrData)
	numChunks := int(math.Ceil(float64(regionInfoLength) / float64(chunkSize)))
	encryptedRegionInfo := make([]byte, 0)
	for i := 0; i < numChunks; i++ {
		from := i * chunkSize
		to := int(math.Min(float64((i+1)*chunkSize), float64(regionInfoLength)))
		chunk := regionCurrData[from:to]
		pubKey, err := endec.RsaParsePubKeyByPrivKey(encPubPrivKey)
		if err != nil {
			logger.Error("parse rsa pub key error: %v", err)
			rspError()
			return
		}
		privKey, err := endec.RsaParsePrivKey(encPubPrivKey)
		if err != nil {
			logger.Error("parse rsa priv key error: %v", err)
			rspError()
			return
		}
		encrypt, err := endec.RsaEncrypt(chunk, pubKey)
		if err != nil {
			logger.Error("rsa enc error: %v", err)
			rspError()
			return
		}
		decrypt, err := endec.RsaDecrypt(encrypt, privKey)
		if err != nil {
			logger.Error("rsa dec error: %v", err)
			rspError()
			return
		}
		if bytes.Compare(decrypt, chunk) != 0 {
			logger.Error("rsa dec test fail")
			rspError()
			return
		}
		encryptedRegionInfo = append(encryptedRegionInfo, encrypt...)
	}
	signPrivkey, err := endec.RsaParsePrivKey(c.signRsaKey)
	if err != nil {
		logger.Error("parse rsa priv key error: %v", err)
		rspError()
		return
	}
	signData, err := endec.RsaSign(regionCurrData, signPrivkey)
	if err != nil {
		logger.Error("rsa sign error: %v", err)
		rspError()
		return
	}
	ok, err := endec.RsaVerify(regionCurrData, signData, &signPrivkey.PublicKey)
	if err != nil {
		logger.Error("rsa verify error: %v", err)
		rspError()
		return
	}
	if !ok {
		logger.Error("rsa verify test fail")
		rspError()
		return
	}
	rsp := &httpapi.QueryCurRegionRspJson{
		Content: base64.StdEncoding.EncodeToString(encryptedRegionInfo),
		Sign:    base64.StdEncoding.EncodeToString(signData),
	}
	context.JSON(http.StatusOK, rsp)
}
