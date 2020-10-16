package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go_backend/config"
	"go_backend/dao/table"
	"go_backend/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	err := initConfig()
	if err != nil {
		log.Error("initConfig error:", err.Error())
	}
	CronDb = InitDb()
}
func InitDb() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=1s&readTimeout=1s", userName, password, ipAddrees, port, dbName, charset)

	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	Db.SetConnMaxLifetime(3600 * time.Second)

	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(50)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&table.Tasks{})
	if err != nil {
		log.Error("AutoMigrate error", err.Error())
	}
	err = db.AutoMigrate(&table.TaskExecHistory{})
	if err != nil {
		panic(err)
	}

	return Db
}
func initConfig() error {
	userNameRes, err := config.GetValue("mysql", "username")
	if err != nil {
		return err
	}
	userName = userNameRes
	passwordRes, err := config.GetValue("mysql", "password")
	if err != nil {
		return err
	}
	password = passwordRes
	ipAddreesRes, err := config.GetValue("mysql", "ipAddrees")
	if err != nil {
		return err
	}
	ipAddrees = ipAddreesRes

	portString, err := config.GetValue("mysql", "port")
	if err != nil {
		return err
	}

	portNew, err := strconv.Atoi(portString)
	if err != nil {
		log.Error("atoi error:%s", err.Error())
		port = -1
	} else {
		port = portNew
	}
	return nil
}
