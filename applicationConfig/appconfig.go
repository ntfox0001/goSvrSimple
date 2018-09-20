package applicationConfig

import (
	"goSvrLib/log"
	"goSvrLib/userSystem"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var (
	Config AppConfig
)

type AppConfig struct {
	UserService userSystem.UserServiceParams `json:"userService"`
}

func InitApplicationConfig(configFilename string) error {
	bytes, err := ioutil.ReadFile(configFilename)
	if err != nil {
		log.Error("config", "ReadFile", err.Error())
		return err
	}
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &Config); err != nil {
		return err
	}
	return nil
}
