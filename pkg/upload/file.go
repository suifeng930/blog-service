package upload

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// 返回经过md5加密之后的 fileName
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext

}

//returns the file name extension
func GetFileExt(name string) string {

	//调用系统方法， 获取file 的后缀名，  ext() ;采用从后往前遍历，遇到 '.' 截取
	return path.Ext(name)
}

// 获取文件保存地址
func GetSavePath() string {

	return global.AppSetting.UploadSavePath

}

// 检查保存目录是否存在
func CheckSavePath(dst string) bool {

	_, err := os.Stat(dst)
	return os.IsNotExist(err)

}

//检查文件后缀名是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {

	ext := GetFileExt(name)

	ext = strings.ToUpper(ext)
	switch TypeImage {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}
	return false

}

//检查文件大小是否超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {

	content, _ := ioutil.ReadAll(f)

	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true

		}
	}
	return false
}

// 检查文件权限是否足够
func CheckPermission(dst string) bool {

	_, err := os.Stat(dst)
	return os.IsPermission(err)

}

// 创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {

	err := os.MkdirAll(dst, perm)
	if err!=nil {
		return err

	}
	return nil
}

// 保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {

	src, err := file.Open()
	if err!=nil {
		return err

	}
	defer src.Close()
	//创建目标地址文件
	out, err := os.Create(dst)
	if err!=nil {
		return err

	}
	defer out.Close()
	// 实现文件copy
	_, err = io.Copy(out, src)
	return err

}