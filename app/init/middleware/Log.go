package middleware

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"net/http"
	"service-all/app/init/Context"
	"service-all/app/init/global"
	"time"
)

func MongoLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		app := Context.Gin{C: c}

		col := global.MONGO.Database("app").Collection("logs")

		_, err := col.InsertOne(c, bson.D{
			{"ip", c.ClientIP()},
			{"method", c.Request.Method},
			{"path", c.FullPath()},
			{"create_at", time.Now().Format("2006-01-02 15:04:05.000")},
		})

		if err != nil {
			global.LOG.Error("写入错误", zap.Error(err))
			app.Response(http.StatusInternalServerError, Context.StatusInternalServerError, nil)
			c.Abort()
		}

		c.Next()
	}
}
