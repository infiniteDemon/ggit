package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"service-all/app/init/Context"
	"service-all/app/init/global"
	"service-all/app/init/models"
)

func AdminValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := Context.Gin{C: c}
		userName, _ := c.Get("username")
		passWord, _ := c.Get("password")
		global.LOG.Info("拦截器", zap.Any("username", userName), zap.Any("password", passWord))

		queryModel := models.ModelAdminUser{}
		act := models.ActAdmin(&queryModel)
		if err := act.QueryWhereUserName(fmt.Sprintf("%s", userName)); err != nil {
			global.LOG.Error("创建管理用户失败", zap.Error(err))
		}
		global.LOG.Debug("用户权限", zap.Any("userid", queryModel.Status))
		if queryModel.Status < 1 {
			global.LOG.Debug("用户权限不足直接驳回", zap.Int("用户权限", queryModel.Status))
			appG.Response(Context.StatusBadRequest, Context.StatusPermissionDenied, nil)
			c.Abort()
		}

		// 处理请求
		c.Next()
	}
}
