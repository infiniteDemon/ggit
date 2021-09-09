package core

import (
	"github.com/spf13/viper"
	"log"
	"service-all/app/init/config"
	"service-all/library/common"
	"service-all/library/file"
	"time"
)

var defaultConf = `{
  "system": {
    "debug": false,
    "host": "127.0.0.1",
    "port": 1111,
    "hash_id": "{HashIDSalt}",
    "hash_session": "{HashIDSalt}",
    "front_web": "/admin"
  },
  "zap": {
    "level": "info",
    "format": "json",
    "author": "DEMON",
    "director": "logs",
    "linkname": "latest_log",
    "encode-level": "LowercaseColorLevelEncoder",
    "stacktrace-key": "stacktrace",
    "log-in-console": true
  },
  "mysql": {
    "is_use": true,
    "user": "root",
    "password": "test123",
    "host": "mysql",
    "port": 3306,
    "db_name": "test",
    "params": "?charset=utf8&parseTime=True&loc=Local",
    "table_prefix": "v1",
    "conn_max_lifetime": 10800,
    "max_open_conns": 10,
    "max_idle_conns": 10
  },
  "redis": {
    "is_use": true,
    "password": "test123",
    "host": "redis",
    "port": 6379,
    "db_name": ""
  },
  "mongo": {
    "is_use": true,
    "user": "root",
    "password": "test123",
    "host": "mongobd",
    "port": 27017,
    "db_name": "",
    "params": "?w=majority",
    "table_prefix": "v1",
    "conn_max_lifetime": 10800,
    "max_open_conns": 10,
    "max_idle_conns": 10
  },
  "wechat": {
    "appid": "{HashIDSalt}",
    "app_secret": "{HashIDSalt}",
    "token": "{HashIDSalt}",
    "encoding_aes_key": "{HashIDSalt}"
  }
}`

func InitConf(path string) *viper.Viper {
	if path == "" || !file.Exists(path) {
		log.Printf("文件不存在，初始化创建基础配置文件")
		// 创建初始配置文件
		confContent := common.Replace(map[string]string{
			"{HashIDSalt}":    common.RandStringRunes(64),
			"{SessionSecret}": common.RandStringRunes(64),
		}, defaultConf)
		f, err := file.CreatNestedFile(path)
		if err != nil {
			log.Panic("无法创建配置文件, %s", err)
		}

		// 写入配置文件
		_, err = f.WriteString(confContent)
		if err != nil {
			log.Panic("无法写入配置文件, %s", err)
		}

		f.Close()
	}
	// alternatively, you can create a new viper instance.
	runtime_viper := viper.New()
	runtime_viper.SetConfigFile(path)

	time.Sleep(time.Second * 5)

	log.Printf("读取配置文件")

	runtime_viper.SetConfigFile(path)
	if err := runtime_viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Panic("Fatal error config file: %s \n", err)
	}

	// unmarshal config
	if err := runtime_viper.Unmarshal(&config.Config); err != nil {
		log.Panic("json unmarshal error: %s", err)
	}
	return runtime_viper
}
