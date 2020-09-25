package service

import (
	"go_backend/config"
	"go_backend/log"
	"os"
	"os/user"
)

var FileSaveDir string

func init() {
	fileDir, err := config.GetValue("share_file", "fileSaveDir")
	if err != nil {
		log.Warn("could not find fileSaveDir in share_file section")
		useHomeDir()
		return
	}
	if !IsExist(FileSaveDir) {
		useHomeDir()
		return
	}
	FileSaveDir = fileDir

}
func useHomeDir() {
	user, err := user.Current()
	if nil != err {
		panic(err)
	}
	FileSaveDir = user.HomeDir
	log.Info("use the home dir:%s", FileSaveDir)
	return
}

// 判断文件是否存在
func IsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		log.Error("IsExist error:%s", err.Error())
		if os.IsExist(err) { // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}
