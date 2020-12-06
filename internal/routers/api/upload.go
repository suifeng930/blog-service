package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}

}

func (u Upload)UploadFile(c *gin.Context)  {

	response := app.NewResponse(c)
	//1.获取到文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	//2。获取到传递过来到文件类型
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err!=nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	if fileHeader ==nil ||fileType<=0{
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	// 3。调用文件处理函数，执行文件保存
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err!=nil {
		global.Logger.Errorf("svc.UploadFile Err:%v",err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{"file_access_url":fileInfo.AccessUrl})

}