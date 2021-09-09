package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author demon
 * @Description //TODO
 * @Date
 * @Param {test string}
 * @return {test string}
 **/

func OptionOk() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.JSON(http.StatusNoContent, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
