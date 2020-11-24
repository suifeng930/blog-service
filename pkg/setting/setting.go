package setting

import (
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {

	vp := viper.New()
	vp.SetConfigName("config")   //指定读取文件名
	vp.AddConfigPath("configs/") //指定读取路径名
	vp.SetConfigType("yaml")     //指定读取文件格式
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("---->",err)
		return nil, err

	}
	return &Setting{vp}, nil
}
