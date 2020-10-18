package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFile(pathStr string, folderName string) (bool, string) {
	folderPath := filepath.Join(pathStr, folderName)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return false, ""
	} else {
		return true, folderPath
	}
}
func GetFilesFromDir(dirPath string) ([]string, error) {

	dir, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}
	fileNameList := make([]string, 0)
	for _, item := range dir {
		fileNameList = append(fileNameList, item.Name())

	}
	return fileNameList, err
}
