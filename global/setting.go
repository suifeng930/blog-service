package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DataBaseSetting *setting.DatabaseSettingS
	DBEngine	*gorm.DB
)