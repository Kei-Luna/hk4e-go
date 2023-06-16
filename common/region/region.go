package region

import (
	"encoding/json"
	"os"

	"hk4e/common/config"
	"hk4e/pkg/endec"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/proto"
)

func LoadRsaKey() (signRsaKey []byte, encRsaKeyMap map[string][]byte, pwdRsaKey []byte) {
	var err error = nil
	encRsaKeyMap = make(map[string][]byte)
	signRsaKey, err = os.ReadFile("key/region_sign_key.pem")
	if err != nil {
		logger.Error("open region_sign_key.pem error: %v", err)
		return nil, nil, nil
	}
	encKeyIdList := []string{"1", "2", "3", "4", "5"}
	for _, v := range encKeyIdList {
		encRsaKeyMap[v], err = os.ReadFile("key/region_enc_key_" + v + ".pem")
		if err != nil {
			logger.Error("open region_enc_key_"+v+".pem error: %v", err)
			return nil, nil, nil
		}
	}
	pwdRsaKey, err = os.ReadFile("key/account_password_key.pem")
	if err != nil {
		logger.Error("open account_password_key.pem error: %v", err)
		return nil, nil, nil
	}
	return signRsaKey, encRsaKeyMap, pwdRsaKey
}

func NewRegionEc2b() *random.Ec2b {
	return random.NewEc2b()
}

// RegionCustomConfig 区服相关的配置，避免在http中使用Json格式
type RegionCustomConfig struct {
	CloseAntiDebug   bool `json:"close_antidebug"`  // 默认false: 默认打开反调开关
	ForceKill        bool `json:"force_kill"`       // 默认false:
	AntiDebugPc      bool `json:"antidebug_pc"`     // 默认false: pc默认不开启反调
	AntiDebugIos     bool `json:"antidubug_ios"`    // 默认false: ios默认不开启反调
	AntiDebugAndroid bool `json:"antidubug_androd"` // 默认false: android默认不开启反调
}

// ClientCustomConfig 客户端版本定义的配置 客户端版本号对应的配置, 需要兼容老的json格式
type ClientCustomConfig struct {
	Visitor        bool              `json:"visitor"`        // 游客功能
	SdkEnv         string            `json:"sdkenv"`         // sdk环境类型
	DebugMenu      bool              `json:"debugmenu"`      // debug菜单
	DebugLogSwitch []int32           `json:"debuglogswitch"` // 打开的log类型
	DebugLog       bool              `json:"debuglog"`       // log总开关
	DeviceList     map[string]string `json:"devicelist"`
	LoadJsonData   bool              `json:"loadjsondata"`  // 用json读取InLevel数据
	ShowException  bool              `json:"showexception"` // 是否显示异常提示框 默认为：true
	CheckDevice    bool              `json:"checkdevice"`
	LoadPatch      bool              `json:"loadPatch"`
	RegionConfig   string            `json:"regionConfig"`
	DownloadMode   int32             `json:"downloadMode"`
}

func GetRegionList(ec2b *random.Ec2b) *proto.QueryRegionListHttpRsp {
	// RegionList
	regionList := new(proto.QueryRegionListHttpRsp)
	regionList.Retcode = 0
	serverList := make([]*proto.RegionSimpleInfo, 0)
	server := &proto.RegionSimpleInfo{
		Name:        "os_usa",
		Title:       "America",
		Type:        "DEV_PUBLIC",
		DispatchUrl: config.GetConfig().Hk4e.DispatchUrl,
	}
	serverList = append(serverList, server)
	regionList.RegionList = serverList
	dispatchEc2bData := ec2b.Bytes()
	regionList.ClientSecretKey = dispatchEc2bData // 客户端使用密钥
	dispatchXorKey := ec2b.XorKey()
	clientCustomConfig, _ := json.Marshal(&ClientCustomConfig{
		SdkEnv:         "2",
		CheckDevice:    false,
		LoadPatch:      false,
		ShowException:  true,
		RegionConfig:   "pm|fk|add",
		DownloadMode:   0,
		DebugMenu:      true,
		DebugLogSwitch: []int32{0, 1, 2},
		DebugLog:       true,
	})
	endec.Xor(clientCustomConfig, dispatchXorKey)
	regionList.ClientCustomConfigEncrypted = clientCustomConfig // 加密后的客户端版本定义的配置
	regionList.EnableLoginPc = true
	return regionList
}

func GetRegionCurr(kcpAddr string, kcpPort int32, ec2b *random.Ec2b) *proto.QueryCurrRegionHttpRsp {
	// RegionCurr
	// region_info: retcode == 0 || RET_STOP_SERVER
	// force_udpate: retcode == RET_CLIENT_FORCE_UPDATE
	// stop_server: retcode == RET_STOP_SERVER
	regionCurr := new(proto.QueryCurrRegionHttpRsp)
	regionCurr.Retcode = 0 // 错误码
	regionCurr.Msg = ""    // 错误信息
	dispatchEc2bData := ec2b.Bytes()
	regionCurr.RegionInfo = &proto.RegionInfo{
		GateserverIp:   kcpAddr,
		GateserverPort: uint32(kcpPort),
		SecretKey:      dispatchEc2bData, // 第一条协议加密密钥
	}
	regionCurr.ClientSecretKey = dispatchEc2bData // 客户端使用密钥
	dispatchXorKey := ec2b.XorKey()
	regionCustomConfig, _ := json.Marshal(&RegionCustomConfig{
		CloseAntiDebug:   true,
		ForceKill:        false,
		AntiDebugPc:      false,
		AntiDebugIos:     false,
		AntiDebugAndroid: false,
	})
	endec.Xor(regionCustomConfig, dispatchXorKey)
	regionCurr.RegionCustomConfigEncrypted = regionCustomConfig // 加密后的区服定义的配置
	clientCustomConfig, _ := json.Marshal(&ClientCustomConfig{
		SdkEnv:         "2",
		CheckDevice:    false,
		LoadPatch:      false,
		ShowException:  true,
		RegionConfig:   "pm|fk|add",
		DownloadMode:   0,
		DebugMenu:      true,
		DebugLogSwitch: []int32{0, 1, 2},
		DebugLog:       true,
	})
	endec.Xor(clientCustomConfig, dispatchXorKey)
	regionCurr.ClientRegionCustomConfigEncrypted = clientCustomConfig // 加密后的客户端区服定义的配置
	return regionCurr
}
