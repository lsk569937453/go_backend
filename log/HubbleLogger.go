package log

var hubbleListLogger = NewLoggerWithonlyMes("hubble.log")

func init() {
	createLogDir()

}

// Info
func Hubbleinfo(mes string) {
	hubbleListLogger.Info(mes)

}

// Error
func HubbleError(mes string) {

	hubbleListLogger.Error(mes)
}
