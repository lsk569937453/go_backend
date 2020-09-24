package config

import (
	"github.com/Unknwon/goconfig"
	"go_backend/log"
)

var config *goconfig.ConfigFile

func init() {

	configNew, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Error("config utils init error:%s", err.Error())
		configNew, err = goconfig.LoadConfigFile("../conf/conf.ini")
		if err != nil {
			panic(err)
		}
	}
	config = configNew
}
func GetValue(section string, key string) (string, error) {
	value, err := config.GetValue(section, key)
	if err != nil {
		log.Error("GetValue error:%s", err.Error())
		return "", err
	} else {
		return value, nil
	}
}
