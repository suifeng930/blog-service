package service

import (
	"errors"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service)UploadFile(fileType upload.FileType,file multipart.File,fileHeader *multipart.FileHeader	) (*FileInfo,error)  {

	//1.获取文件名
	fileName := upload.GetFileName(fileHeader.Filename)
	//2。获取文件写入路径
	uploadSavePath := upload.GetSavePath()
	//组装文件写入路径
	dst :=uploadSavePath+"/"+fileName

	//3.判断文件后村名是否匹配
	if !upload.CheckContainExt(fileType,fileName) {
		return nil,errors.New("file suffix is not supported")

	}

	//4。判断文件路径是否匹配
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err!=nil {
			return nil, errors.New("failed to create save directory")

		}

	}
	//5。判断文件大小是否匹配
	if upload.CheckMaxSize(fileType,file) {
		return nil,errors.New("exceeded maxinum file limit ")

	}
	//6。判断操作文件权限是否匹配
	if upload.CheckPermission(uploadSavePath) {
		return nil,errors.New("insufficient file permissions ")

	}
	//7。执行写入文件操作
	if err:=upload.SaveFile(fileHeader,dst);err!=nil {
		return nil, err

	}
	accessUrl :=global.AppSetting.UploadServerUrl+"/"+fileName
	return &FileInfo{Name: fileName,AccessUrl: accessUrl},nil
}