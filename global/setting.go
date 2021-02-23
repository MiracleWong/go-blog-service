package global

import (
	"github.com/MiracleWong/go-blog-service/pkg/logger"
	"github.com/MiracleWong/go-blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	DataSetting *setting.DatabaseSettingS
	AppSetting *setting.AppSettingS
	// Logger 全局变量
	Logger *logger.Logger
)

