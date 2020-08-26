package dao

import (
	"fmt"
	"go_backend/config"
	"go_backend/log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (


	userName  string = ""
	password  string = ""
	ipAddrees string = ""
	port      int    = -1
	dbName    string = "cron_timer"
	charset   string = "utf8"
)
var CronDb *sqlx.DB

func init() {
	initConfig()
	CronDb = InitDb()
}
func InitDb() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}
func initConfig() {
	userName = config.GetValue("mysql", "username")
	password = config.GetValue("mysql", "password")
	ipAddrees = config.GetValue("mysql", "ipAddrees")
	portString := config.GetValue("mysql", "port")
	portNew, err := strconv.Atoi(portString)
	if err != nil {
		log.Error("atoi error:%s", err.Error())
		port = -1
	} else {
		port = portNew
	}
}
