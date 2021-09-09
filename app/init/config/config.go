package config

import (
	"time"
)

type system struct {
	Debug       bool   `mapstructure:"debug" yaml:"debug" json:"debug"`
	Host        string `mapstructure:"host" yaml:"host" json:"host"`
	Port        int    `mapstructure:"port" yaml:"port" json:"port"`
	JwtSecret   string `mapstructure:"jwt_secret" yaml:"jwt_secret" json:"jwt_secret"`
	Issuer      string `mapstructure:"issuer" yaml:"issuer" json:"issuer"`
	HashId      string `mapstructure:"hash_id" yaml:"hash_id" json:"hash_id"`
	HashSession string `mapstructure:"hash_session" yaml:"hash_session" json:"hash_session"`
	FrontWeb    string `mapstructure:"front_web" yaml:"front_web" json:"front_web"`
}

type obs struct {
	Ws     string `mapstructure:"ws" yaml:"ws" json:"ws"`
	Origin string `mapstructure:"origin" yaml:"origin" json:"origin"`
}

type zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Author        string `mapstructure:"author" json:"author" yaml:"author"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

type mysql struct {
	IsUse           bool          `mapstructure:"is_use" yaml:"is_use" json:"is_use"`
	User            string        `mapstructure:"user" yaml:"user" json:"user"`
	Password        string        `mapstructure:"password" yaml:"password" json:"password"`
	Host            string        `mapstructure:"host" yaml:"host" json:"host"`
	Port            int           `mapstructure:"port" yaml:"port" json:"port"`
	DbName          string        `mapstructure:"db_name" yaml:"db_name" json:"db_name"`
	Params          string        `mapstructure:"params" yaml:"params" json:"params"`
	TablePrefix     string        `mapstructure:"table_prefix" yaml:"table_prefix" json:"table_prefix"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime" json:"conn_max_lifetime"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
}

type backup struct {
	IsUse           bool          `mapstructure:"is_use" yaml:"is_use" json:"is_use"`
	User            string        `mapstructure:"user" yaml:"user" json:"user"`
	Password        string        `mapstructure:"password" yaml:"password" json:"password"`
	Host            string        `mapstructure:"host" yaml:"host" json:"host"`
	Port            int           `mapstructure:"port" yaml:"port" json:"port"`
	DbName          string        `mapstructure:"db_name" yaml:"db_name" json:"db_name"`
	Params          string        `mapstructure:"params" yaml:"params" json:"params"`
	TablePrefix     string        `mapstructure:"table_prefix" yaml:"table_prefix" json:"table_prefix"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime" json:"conn_max_lifetime"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
}

type mongo struct {
	IsUse           bool          `mapstructure:"is_use" yaml:"is_use" json:"is_use"`
	User            string        `mapstructure:"user" yaml:"user" json:"user"`
	Password        string        `mapstructure:"password" yaml:"password" json:"password"`
	Host            string        `mapstructure:"host" yaml:"host" json:"host"`
	Port            int           `mapstructure:"port" yaml:"port" json:"port"`
	DbName          string        `mapstructure:"db_name" yaml:"db_name" json:"db_name"`
	Params          string        `mapstructure:"params" yaml:"params" json:"params"`
	TablePrefix     string        `mapstructure:"table_prefix" yaml:"table_prefix" json:"table_prefix"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime" json:"conn_max_lifetime"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
}

type redis struct {
	IsUse    bool   `mapstructure:"is_use" yaml:"is_use" json:"is_use"`
	Password string `mapstructure:"password" yaml:"password" json:"password"`
	Host     string `mapstructure:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" yaml:"port" json:"port"`
	DbName   string `mapstructure:"db_name" yaml:"db_name" json:"db_name"`
}

type wechat struct {
	Appid          string `mapstructure:"appid" yaml:"appid" json:"appid"`
	AppSecret      string `mapstructure:"app_secret" yaml:"app_secret" json:"app_secret"`
	Token          string `mapstructure:"token" yaml:"token" json:"token"`
	EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key" json:"encoding_aes_key"`
}

type aliyun struct {
	OssHost         string `json:"oss_host" mapstructure:"oss_host" yaml:"oss_host"`
	AccessKeyId     string `json:"access_key_id" mapstructure:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"access_key_secret" yaml:"access_key_secret"`
	Host            string `json:"host" mapstructure:"host" yaml:"host"`
	UploadDir       string `json:"upload_dir" mapstructure:"upload_dir" yaml:"upload_dir"`
	CallbackUrl     string `json:"callback_url" mapstructure:"callback_url" yaml:"callback_url"`
	ExpireTime      int64  `json:"expire_time" mapstructure:"expire_time" yaml:"expire_time"`
}

type smallWechatApp struct {
	Appid          string `mapstructure:"appid" yaml:"appid" json:"appid"`
	AppSecret      string `mapstructure:"app_secret" yaml:"app_secret" json:"app_secret"`
	Token          string `mapstructure:"token" yaml:"token" json:"token"`
	EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key" json:"encoding_aes_key"`
}

type config struct {
	System         system
	Obs            obs
	Zap            zap
	Mysql          mysql
	Backup         backup
	Mongo          mongo
	Redis          redis
	Wechat         wechat
	Aliyun         aliyun
	SmallWechatApp smallWechatApp
}

var Config = &config{}
