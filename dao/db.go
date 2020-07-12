package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

var (
	userName  string = "root"
	password  string = "elong"
	ipAddrees string = "10.160.85.246"
	port      int    = 3308
	dbName    string = "cron_timer"
	charset   string = "utf8"
)
var CronDb *sqlx.DB

func init() {

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
