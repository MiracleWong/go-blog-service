package global

import (
	"github.com/MiracleWong/go-blog-service/pkg/setting"
	"gorm.io/gorm"
)

// 声明全局变量
var (
	ServerSetting *setting.ServerSettingS
	DataSetting   *setting.DatabaseSettingS
	AppSetting    *setting.AppSettingS
	DBEngine      *gorm.DB
)
