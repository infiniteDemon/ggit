package middleware

import (
	"github.com/gin-gonic/gin"
	"service-all/library/logger"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//app := Context.Gin{C: c}
		email, _ := c.Get("email")
		msg, _ := c.Get("msg")
		id, _ := c.Get("id")

		act := c.Request.Method
		domain := c.Param("domain")
		pid := c.Param("pid")

		logger.Log().Info("id 为 %v， email 为 %s的用户，想要访问%s域中的项目id为%s的项目，携带了一个msg为%s, 动作为%s", id, email, domain, pid, msg, act)

		//res, err := rbca.InitRbca().Enforce(email, domain, pid, act)
		//if res {
		//	// hh default data3 read
		//	logger.Log().Info("权限验证通过 %v", res)
		//	c.Next()
		//} else {
		//	logger.Log().Error("无访问权限，%v %s", res, err)
		//	app.Response(http.StatusUnauthorized, Context.StatusUnauthorized, err)
		//	c.Abort()
		//}
		c.Next()
	}
}
