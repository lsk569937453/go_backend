package log

var heartBeat = NewCommonlogger("heartBeat.log")

func init() {
	createLogDir()
}

// Info
func HeartBeatinfo(args string) {
	heartBeat.Info(args)
}

// Error
func HeartBeatError(template string) {
	heartBeat.Error(template)
}
