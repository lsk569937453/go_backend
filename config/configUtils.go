package config

import (
	"github.com/Unknwon/goconfig"
	"go_backend/log"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

var config *goconfig.ConfigFile

func init() {

	configNew, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Error("config utils init error:%s", err.Error())
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		for {
			if !isContainsConf(dir) {
				dir = GetParentDirectory(dir)
			} else {
				break
			}
		}
		configNew, err = goconfig.LoadConfigFile(dir + "/conf/conf.ini")
		if err != nil {
			panic(err)
		}
	}
	config = configNew
}
func isContainsConf(pathStr string) bool {
	rd, err := ioutil.ReadDir(pathStr)

	if err != nil {
		log.Error("isContainsConf error:", err)
	}

	for _, fi := range rd {
		if fi.IsDir() && fi.Name() == "conf" {
			return true
		}

	}
	return false

}
func GetParentDirectory(dirctory string) string {

	dirChar := "/"
	if runtime.GOOS == "windows" {
		dirChar = "\\"
	}

	return substr(dirctory, 0, strings.LastIndex(dirctory, dirChar))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
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
