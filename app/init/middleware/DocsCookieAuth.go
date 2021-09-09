package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"service-all/app/controller/CDocsFun"
	"service-all/app/init/Context"
	"service-all/app/init/global"
)

// docs
func DocsCookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("gin_docs_cookie")

		if err != nil {
			cookie = "NotSet"
			c.Redirect(http.StatusFound, "/docs/v1/login")
		}
		data, _ := CDocsFun.ParseToken(cookie)
		global.LOG.Debug("Cookie value", zap.String("data", cookie), zap.Any("解析参数", data))
		DocsAuthUser(data.UserName, data.PassWord, c)
		// 处理请求
		c.Next()
	}
}

func DocsAuthUser(UserName, Password string, c *gin.Context) {
	app := Context.Gin{C: c}
	if UserName != "nizi" {
		global.LOG.Error("账户错误")
		app.Response(http.StatusBadRequest, Context.StatusBadRequest, "账户错误")
		c.Abort()
		return
	}

	if Password != "ef914377f626b4de0be35759b6625e3b" {
		global.LOG.Error("密码错误")
		app.Response(http.StatusBadRequest, Context.StatusBadRequest, "密码错误")
		c.Abort()
		return
	}

	global.LOG.Debug("文档中间件过滤判断成功")
}
