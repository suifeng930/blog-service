package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {

	vp := viper.New()
	vp.SetConfigName("config") //指定读取文件名
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)

		}
	}
	//vp.AddConfigPath("configs/") //指定读取路径名
	vp.SetConfigType("yaml") //指定读取文件格式
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("---->", err)
		return nil, err

	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

//
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()

}
