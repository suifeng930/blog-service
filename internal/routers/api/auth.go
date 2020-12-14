package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

// 获取鉴权信息
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errors)
		errResp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errResp)
		return

	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthrizedAuthNotExist)
		return

	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})

}
