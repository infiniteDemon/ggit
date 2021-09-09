package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"service-all/app/init/Context"
)

/**
 * @Author demon
 * @Description //TODO
 * @Date
 * @Param {test string}
 * @return {test string}
 **/
// JWT is jwt middleware

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := Context.Gin{C: c}
		token := c.Request.Header.Get("Authorization")
		Len := len(token)
		if Len <= 0 || token == "" {
			appG.Response(Context.StatusUnauthorized, Context.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		token = token[7:Len]
		parseToken, err := ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				appG.Response(Context.StatusUnauthorized, Context.StatusUnauthorized, nil)
				c.Abort()
				return
			case jwt.ValidationErrorIssuer:
				appG.Response(Context.StatusUnauthorized, Context.StatusUnauthorized, nil)
				c.Abort()
				return
			default:
				appG.Response(Context.StatusUnauthorized, Context.StatusUnauthorized, nil)
				c.Abort()
				return
			}
		}
		c.Set("username", parseToken.UserName)
		c.Set("password", parseToken.PassWord)
		c.Next()
	}
}
