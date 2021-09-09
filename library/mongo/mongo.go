package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"service-all/app/init/config"
	"service-all/app/init/global"
	"time"
)

func InitMongo() *mongo.Client {

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s%s",
		config.Config.Mongo.User,
		config.Config.Mongo.Password,
		config.Config.Mongo.Host,
		config.Config.Mongo.Port,
		config.Config.Mongo.DbName,
		config.Config.Mongo.Params,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Config.Mongo.ConnMaxLifetime)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		global.LOG.Error("mongodb连接失败", zap.Error(err))
	}

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		global.LOG.Error("mongodb解除连接失败", zap.Error(err))
	//	}
	//}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	//
	global.LOG.Info("mongodb测试连接成功")
	return client
}
