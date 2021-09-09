package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func BindAndCheck(ctx *gin.Context, data interface{}) error {
	if err := ctx.ShouldBindJSON(data); err != nil {
		return errors.New(fmt.Sprintf("bindjson err%s", err))
	}
	// 校验数据
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return errors.New(fmt.Sprintf("validator err%s", err))
	}
	return nil
}
