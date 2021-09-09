package orm

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"service-all/app/init/config"
	"service-all/app/init/global"
	"strings"
	"time"
)

/*
**
* @Author demon
* @Description //TODO gorm模块连接数据库
* @Date 2020-7-12 20:03:07
**/

// Init 初始化 MySQL 链接
func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v%v",
		config.Config.Mysql.User,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.DbName,
		config.Config.Mysql.Params,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 彩色打印
		},
	)
	if config.Config.System.Debug {
		newLogger.LogMode(logger.Silent)
	} else {
		newLogger.LogMode(logger.Info)
	}

	//// 处理表前缀

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s_", config.Config.Mysql.TablePrefix), // 表名前缀，`User`表为`t_users`
			SingularTable: true,                                                // 使用单数表名，启用该选项后，`User` 表将是`user`
			NameReplacer:  strings.NewReplacer("/", "_"),                       // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})

	if err != nil {
		global.LOG.Panic("数据库连接失败", zap.Error(err))
	}

	sqlDB, sqlErr := db.DB()
	sqlDB.SetConnMaxLifetime(time.Second * config.Config.Mysql.ConnMaxLifetime)
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConns)

	if sqlErr != nil {
		global.LOG.Panic("数据库通用配置失败", zap.Error(err))
	}

	// Ping
	if pingErr := sqlDB.Ping(); pingErr != nil {
		global.LOG.Panic("数据库测试连接失败", zap.Error(pingErr))
	}

	global.LOG.Info("MYSQL数据库测试连接成功")
	return db
}
