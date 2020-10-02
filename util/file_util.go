package util

import (
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
