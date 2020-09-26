package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go_backend/log"
	"go_backend/redis"
	"go_backend/util"
	"go_backend/vojo"
	"mime/multipart"
	"net/http"
	"net/url"
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
func SaveFile(c *gin.Context, fStream *multipart.FileHeader) (string, error) {
	newFileName := util.GetCurrentTime() + DefaultKeyCombineChar + fStream.Filename
	dst := filepath.Join(FileSaveDir, newFileName)
	err := c.SaveUploadedFile(fStream, dst)

	var realKey string
	if err != nil {
		log.Error("bind error:%v", err.Error())
		return "", err

	} else {
		uu := uuid.NewV4()
		realKey = uu.String()[:8]
		err = saveExpireTime(realKey, newFileName)
		if err != nil {
			log.Error("saveExpireTime err:%s", err.Error())
			return "", err
		}
		err = saveCount(realKey, DefaultDownloadCount)
		if err != nil {
			log.Error("saveCount err:%s", err.Error())
			return "", err
		}

	}
	return realKey, nil
}
func DownloadService(ctx *gin.Context) error {
	var req vojo.DownloadFileReq

	//fileName := ctx.Param("fileKey")
	err := ctx.BindJSON(&req)
	if err != nil {
		return err
	}
	fileName := req.FileKeyCode
	exist, realName := isFileExpire(fileName)
	if !exist {
		return errors.New("file not exits")
	}
	if !isFileCountLegal(fileName) {
		return errors.New("file count is 0")
	}

	targetPath := filepath.Join(FileSaveDir, realName)

	fileLen := len(realName)
	fileStart := len(util.TimeFormat + DefaultKeyCombineChar)
	fileName = realName[fileStart:fileLen]
	fileName = url.QueryEscape(fileName)

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("share-file-name", fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
	ctx.Status(http.StatusOK)

	log.Info("%s has down load over ", realName)
	return nil
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
func isFileExpire(fileKey string) (bool, string) {
	redisFileKey := DefaultExKeyPrefix + DefaultKeyCombineChar + fileKey
	duration, err := redis.TTL(redisFileKey)
	if err != nil {
		log.Error("TTL %s", err.Error())
	}
	if duration > 0 {
		fileNameWithTime, err := redis.Get(redisFileKey)
		if err != nil {
			return true, ""
		} else {

			return true, fileNameWithTime
		}
	} else {
		return false, ""
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
