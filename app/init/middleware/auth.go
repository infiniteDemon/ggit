package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//app := Context.Gin{C: c}
		if len(c.GetHeader("Authorization")) == 0 {
			//app.Response(http.StatusUnauthorized, Context.StatusUnauthorized, "Authorization is required Header")
			//c.Abort()
			c.Redirect(http.StatusFound, "/web/user/login")
		}
		c.Next()
	}
}
