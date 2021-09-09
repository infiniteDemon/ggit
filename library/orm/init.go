package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

/*
**
* @Author demon
* @Description //TODO gorm模块连接数据库
* @Date 2020-7-12 20:03:07
**/

// DB 数据库链接单例
var DB *gorm.DB

// Init 初始化 MySQL 链接
func Init(User, Password, Host, Port, DbName, TablePrefix, Type string, Debug bool) {
	log.Printf("初始化数据库连接")

	var (
		db  *gorm.DB
		err error
	)

	str := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DbName)

	db, err = gorm.Open(Type, str)

	// 处理表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return TablePrefix + defaultTableName
	}

	// Debug模式下，输出所有 SQL 日志
	if Debug {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if err != nil {
		log.Printf("连接数据库不成功, %s", err)
		return
	}

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(10)
	//打开
	db.DB().SetMaxOpenConns(20)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 300)

	DB = db
}
