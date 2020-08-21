package log

var thriftLogger = NewCommonlogger("thrift.log")

func init() {
	createLogDir()
}

// Info
func Thriftinfo(args string) {
	thriftLogger.Info(args)
}

// Error
func ThriftError(template string) {
	thriftLogger.Error(template)
}
