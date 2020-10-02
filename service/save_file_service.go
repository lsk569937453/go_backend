package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_backend/log"
	"go_backend/redis"
	"go_backend/util"
	"go_backend/vojo"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	DefaultDownloadCount  = 1
	DefaultExKeyPrefix    = "file"
	DefaultCountKeyPrefix = "count"
	DefaultKeyCombineChar = "=="
)

/**
 *
 * @Description saveRedisFile
 * @Date 3:36 下午 2020/9/25
 **/
func SaveFile(c *gin.Context, fStream *multipart.FileHeader, clientId string) (string, error) {
	newFileName := util.GetCurrentTime() + DefaultKeyCombineChar + fStream.Filename
	//	dst := filepath.Join(FileSaveDir, newFileName)
	isCreate, dst := util.CreateFile(FileSaveDir, clientId)
	if !isCreate {
		err := fmt.Errorf("CreateFile error")
		log.Error("SaveFile error ", err)
		return "", err
	}
	dst = filepath.Join(dst, newFileName)

	err := c.SaveUploadedFile(fStream, dst)

	if err != nil {
		log.Error("bind error:%v", err.Error())
		return "", err

	} else {

		err = saveFileTime(clientId, newFileName)
		if err != nil {
			log.Error("saveExpireTime err:%s", err.Error())
			return "", err
		}
		err = saveCount(clientId, DefaultDownloadCount)
		if err != nil {
			log.Error("saveCount err:%s", err.Error())
			return "", err
		}

	}
	return clientId, nil
}
func DownloadService(ctx *gin.Context) error {
	var req vojo.DownloadFileReq

	//fileName := ctx.Param("fileKey")
	err := ctx.BindJSON(&req)
	if err != nil {
		return err
	}
	clientId := req.FileKeyCode
	frontEndFileName := req.FileName
	exist, realName := isFileExpire(clientId, frontEndFileName)
	if !exist {
		return errors.New("file not exits")
	}
	if !isFileCountLegal(clientId) {
		return errors.New("file count is 0")
	}

	targetPath := filepath.Join(FileSaveDir, clientId, realName)

	fileLen := len(realName)
	fileStart := len(util.TimeFormat + DefaultKeyCombineChar)
	fileName := realName[fileStart:fileLen]
	fileName = url.QueryEscape(fileName)

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("share-file-name", fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	file, err := os.Open(targetPath)
	if err != nil {
		log.Error("open error:", err)
		return err
	}
	defer file.Close()

	io.Copy(ctx.Writer, file)
	ctx.Status(http.StatusOK)

	log.Info("%s has down load over ", realName)

	return nil
}

func GetFileList(clientId string) ([]string, error) {
	res, err := redis.HGetAll(DefaultExKeyPrefix + DefaultKeyCombineChar + clientId)

	if err != nil {
		return nil, err
	}
	fileList := make([]string, 0)
	for key, _ := range res {
		fileLen := len(key)
		fileStart := len(util.TimeFormat + DefaultKeyCombineChar)
		fileName := key[fileStart:fileLen]
		fileList = append(fileList, fileName)
	}
	return fileList, err
}

/**
 *
 * @Description
 * @Date 4:32 下午 2020/9/25
 **/
func isFileCountLegal(fileKey string) bool {
	redisCountKey := DefaultCountKeyPrefix + DefaultKeyCombineChar + fileKey
	countString, err := redis.Get(redisCountKey)
	if err != nil {
		log.Error("Get error %s", err.Error())
		return false
	}
	count, err := strconv.Atoi(countString)
	if err != nil {
		log.Error("Atoi error %s", err.Error())
		return false
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

/**
 *
 * @Description  get the file status
 * @Date 4:25 下午 2020/9/25
 **/
func isFileExpire(fileKey string, fileName string) (bool, string) {
	redisFileKey := DefaultExKeyPrefix + DefaultKeyCombineChar + fileKey
	duration, err := redis.TTL(redisFileKey)
	if err != nil {
		log.Error("TTL %s", err.Error())
	}
	if duration > 0 {
		fileNameWithTimeMap, err := redis.HGetAll(redisFileKey)
		if err != nil {
			return false, ""
		} else {
			for key, _ := range fileNameWithTimeMap {
				fileLen := len(key)
				fileStart := len(util.TimeFormat + DefaultKeyCombineChar)
				itemFileName := key[fileStart:fileLen]
				if itemFileName == fileName {
					return true, key
				}
			}

			return false, ""
		}
	} else {
		return true, ""
	}

}

/**
 *
 * @Description
 * @Date 3:37 下午 2020/9/25
 **/
func saveCount(srcKey string, count int) error {
	err := redis.SetNX(DefaultCountKeyPrefix+DefaultKeyCombineChar+srcKey, count, time.Hour*24)
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	return nil

}

/**
 *
 * @Description
 * @Date 3:37 下午 2020/9/25
 **/
func saveExpireTime(srcKey string, newFileName string) error {
	err := redis.SetNX(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, newFileName, time.Hour*24)
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	return nil
}
func saveFileTime(srcKey string, newFileName string) error {
	err := redis.HSet(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, newFileName, "-1")
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	redis.Expire(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, time.Hour*24)
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	return nil
}
