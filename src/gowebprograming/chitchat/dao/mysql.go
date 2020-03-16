package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var DB *sqlx.DB

const (
	user     string = "root"
	password string = "root"
	ip       string = "localhost"
	port     int    = 3306
	dbName   string = "chitchat"
	charset  string = "utf8"
)

func init() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true", user, password, ip, port, dbName, charset))
	if err != nil {
		log.Fatal("连接数据库错误！", err)
		return
	}
	DB = db
	DB.SetConnMaxLifetime(5 * time.Second)
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	log.Println("连接数据库成功！")
}
