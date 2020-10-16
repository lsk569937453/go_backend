package log

import (
	"fmt"
	"os"
	"os/user"
)

var (
	logPath string = createLogDir()
)
var logger = NewCommonlogger("commonGo.log")

func init() {
	createLogDir()
}
func createLogDir() string {
	user, err := user.Current()
	if nil != err {
		fmt.Println("", err.Error())
		return "/home/work/ddalog/"
	}
	fmt.Println("log dir:" + user.HomeDir)
	return user.HomeDir + "/"

	//err := os.MkdirAll(logPath, 0777)
	//if err != nil {
	//	fmt.Printf("%s", err)
	//} else {
	//	fmt.Print("创建目录成功!")
	//}
}
func BaseGinLog() *os.File {
	logfileName := logPath + "gin.access.log"
	logFile, logErr := os.OpenFile(logfileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	return logFile

}

// Debug

func Debug(mes string, args ...interface{}) {
	logger.Sugar().Errorf(mes, args)

}
func Info(mes string, args ...interface{}) {
	logger.Sugar().Infof(mes, args)

}

// Warn
func Warn(args string) {
	logger.Warn(args)
}

// Error
func Errorf(mes string, args ...interface{}) {
	logger.Sugar().Errorf(mes, args)

}

// Error
func Error(mes string, args ...interface{}) {

	logger.Sugar().Errorf(mes, args)

}
