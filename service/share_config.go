package service

import (
	"go_backend/config"
	"go_backend/log"
	"os"
	"os/user"
	"path/filepath"
)

var FileSaveDir string

const Share_File_Dir = "shareFile"

func init() {
	fileDir, err := config.GetValue("share_file", "fileSaveDir")
	if err != nil {
		log.Warn("could not find fileSaveDir in share_file section")
		useHomeDir()
		return
	}
	if !IsExist(fileDir) {
		useHomeDir()
		return
	}
	isCreate, filePath := createFile(fileDir, Share_File_Dir)
	if isCreate {
		FileSaveDir = filePath
	} else {
		panic("create dir error")
	}
}
func createFile(pathStr string, folderName string) (bool, string) {
	folderPath := filepath.Join(pathStr, folderName)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return false, ""
	} else {
		return true, folderPath
	}

}
func useHomeDir() {

	user, err := user.Current()
	if nil != err {
		panic(err)
	}
	isCreate, filePath := createFile(user.HomeDir, Share_File_Dir)
	if isCreate {
		FileSaveDir = filePath
	} else {
		panic("create dir error")
	}
	log.Info("use the home dir:%s", FileSaveDir)
}

// 判断文件是否存在
func IsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		log.Error("IsExist error:%s", err.Error())
		return os.IsExist(err)
	}
	return true
}
