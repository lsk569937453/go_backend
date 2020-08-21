package log

import (
	"fmt"
	"os/user"
)

var (
	logPath string = createLogDir()
)
var logger = NewCommonlogger("dda.log")



func init() {
	createLogDir()
}
func createLogDir() string{
	user, err := user.Current()
	if nil != err {
		fmt.Printf("",err)
		return "/home/work/ddalog/"
	}
	return user.HomeDir+"/"

	//err := os.MkdirAll(logPath, 0777)
	//if err != nil {
	//	fmt.Printf("%s", err)
	//} else {
	//	fmt.Print("创建目录成功!")
	//}
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
