package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string   //文件的最终保存路径
	UploadServerUrl      string   // 上传文件后用于展示的文件服务地址
	UploadImageMaxSize   int      //上传文件所允许的最大空间大小
	UploadImageAllowExts []string //上传文件所允许的文件后缀
}
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err

	}
	return nil
}
