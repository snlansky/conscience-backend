package http

import (
	"conscience-backend/model"
	"conscience-backend/service"
	"github.com/gin-gonic/gin/binding"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const RequestParametersError = "请求参数错误"

func RegisterFaceID(ctx *gin.Context){
	var req model.RequestRegisterFaceID

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	err = service.DefaultFaceIDService.RegisterFaceID(&req)
	if err != nil {
		logger.Errorf("client %s request invoke contact error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(""))
}


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
