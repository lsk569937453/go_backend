package log

var validateSqlLogger = NewCommonlogger("validateSql.log")

func init() {
	createLogDir()

}

// Info
func ValidateSqlinfo(args string) {
	validateSqlLogger.Info(args)
}

// Error
func ValidateSqlError(args string) {
	validateSqlLogger.Error(args)
}
