package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_backend/log"
	"go_backend/redis"
	"go_backend/util"
	"go_backend/vojo"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultDownloadCount  = 1
	DefaultExKeyPrefix    = "file"
	DefaultCountKeyPrefix = "count"
	DefaultKeyCombineChar = "=="
)

func SaveChunk(c *gin.Context) (string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return "", err
	}
	clientId := form.Value["clientId"]
	if len(clientId) == 0 || clientId[0] == "" {
		return "", fmt.Errorf("form value error,clientId is null")
	}
	chunkIndex := form.Value["index"]
	if len(chunkIndex) == 0 || chunkIndex[0] == "" {
		return "", fmt.Errorf("form value error,index is null")
	}
	files := form.File["file"]
	if len(files) == 0 {
		return "", fmt.Errorf("form value error,file is null")
	}
	fileName := form.Value["fileName"]
	if len(fileName) == 0 {
		return "", fmt.Errorf("form value error,fileName is null")
	}

	isCreate, dst := util.CreateFile(FileSaveDir, clientId[0])
	if !isCreate {
		err := fmt.Errorf("CreateFile error")
		log.Error("SaveFile error ", err)
		return "", err
	}

	realFileName := chunkIndex[0] + fileName[0]
	dst = filepath.Join(dst, realFileName)

	err = c.SaveUploadedFile(files[0], dst)

	return clientId[0], err
}

type FileListSort []os.FileInfo

func (s FileListSort) Len() int      { return len(s) }
func (s FileListSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s FileListSort) Less(i, j int) bool {
	return s[i].Name() < s[j].Name()
}

//merge the file chunk

func MergeChunk(c *gin.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	clientId := form.Value["clientId"]
	if len(clientId) == 0 || clientId[0] == "" {
		return fmt.Errorf("form value error,clientId is null")
	}

	fileName := form.Value["fileName"]
	if len(fileName) == 0 {
		return fmt.Errorf("form value error,fileName is null")
	}
	isCreate, dst := util.CreateFile(FileSaveDir, clientId[0])
	if !isCreate {
		return fmt.Errorf("createFile error")
	}
	//get allFiles in the dir
	dirFileList, err := ioutil.ReadDir(dst)
	if err != nil {
		return err
	}
	//filter the chunkFile
	realChunkFileList := make(FileListSort, 0)
	for _, value := range dirFileList {
		everyFileName := value.Name()
		if strings.Contains(everyFileName, fileName[0]) {
			realChunkFileList = append(realChunkFileList, value)
		}

	}
	sort.Stable(realChunkFileList)
	if len(realChunkFileList) == 0 {
		return fmt.Errorf("file chunk count is 0")
	}
	realFileName := filepath.Join(dst, fileName[0])

	out, err := os.OpenFile(realFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer out.Close()
	wt := bufio.NewWriter(out)

	for _, item := range realChunkFileList {
		chunkPath := filepath.Join(dst, item.Name())
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}
		defer chunkFile.Close()

		_, err = io.Copy(wt, chunkFile)
		if err != nil {
			return err
		}
	}
	wt.Flush()

	return nil

}

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
func DownloadChunk(ctx *gin.Context) error {
	var req vojo.DownloadFileReq

	err := ctx.BindJSON(&req)
	if err != nil {
		return err
	}
	clientId := req.FileKeyCode
	frontEndFileName := req.FileName

	targetPath := filepath.Join(FileSaveDir, clientId, frontEndFileName)

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("share-file-name", frontEndFileName)
	ctx.Header("Content-Type", "application/octet-stream")
	file, err := os.Open(targetPath)
	if err != nil {
		log.Error("open error:", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK)

	log.Info("%s has down load over ", frontEndFileName)

	return nil
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
	fileStart := len(util.TimeFormatFirst + DefaultKeyCombineChar)
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

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		return err
	}
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
	for key := range res {
		fileLen := len(key)
		fileStart := len(util.TimeFormatFirst + DefaultKeyCombineChar)
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
			for key := range fileNameWithTimeMap {
				fileLen := len(key)
				fileStart := len(util.TimeFormatFirst + DefaultKeyCombineChar)
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

///**
// *
// * @Description
// * @Date 3:37 下午 2020/9/25
// **/
//func saveExpireTime(srcKey string, newFileName string) error {
//	err := redis.SetNX(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, newFileName, time.Hour*24)
//	if err != nil {
//		log.Error("set redis  error:%v", err.Error())
//		return err
//	}
//	return nil
//}
func saveFileTime(srcKey string, newFileName string) error {
	err := redis.HSet(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, newFileName, "-1")
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	err = redis.Expire(DefaultExKeyPrefix+DefaultKeyCombineChar+srcKey, time.Hour*24)
	if err != nil {
		log.Error("set redis  error:%v", err.Error())
		return err
	}
	return nil
}
