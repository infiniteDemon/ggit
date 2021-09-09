package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

/**
 * @Author demon
 * @Description //TODO mysql连接
 * @Date 2020-7-12 18:00:39
 **/

//Db数据库连接池
var DB *sql.DB

func Connect(User, Password, Host, Port, DbName string) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{User, ":", Password, "@tcp(", Host, ":", Port, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(20)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Println("opon database fail", err)
		return
	}
	log.Println("mysql connnect success")
}
