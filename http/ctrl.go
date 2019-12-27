package http

import (
	"github.com/gin-gonic/gin/binding"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const RequestParametersError = "请求参数错误"

func bind(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBind(obj)
	if err != nil {
		return err
	}

	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}

func bindBody(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindBodyWith(obj, binding.JSON)
	if err != nil {
		return err
	}
	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}

func bindQuery(ctx *gin.Context, obj interface{}) error {
	err := ctx.BindQuery(obj)
	if err != nil {
		return err
	}
	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}
