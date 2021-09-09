package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	BackupDB *gorm.DB
	Conn     *websocket.Conn
	REDIS    *redis.Client
	MONGO    *mongo.Client
	VP       *viper.Viper
	LOG      *zap.Logger
)
