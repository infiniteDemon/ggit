package Context

/**
 * @Author demon
 * @Description //TODO 返回请求基础序列器
 * @Date 2020-7-12 17:25:50
 **/

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// Response 基础序列化器
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := statusText[code]
	if ok {
		return msg
	}

	return statusText[StatusBadRequest]
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	_, ok1 := data.(error)
	if ok1 {
		data = fmt.Sprintf("%s", data)
	}
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPServerError example
type HTTPServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

// HTTPNotFound example
type HTTPNotFound struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Request Not Found"`
}
