package region

import (
	"os"

	"hk4e/pkg/logger"
	"hk4e/pkg/random"
)

func LoadRegionRsaKey() (signRsaKey []byte, encRsaKeyMap map[string][]byte, pwdRsaKey []byte) {
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

func LoadRegionEc2b() *random.Ec2b {
	return random.NewEc2b()
}
