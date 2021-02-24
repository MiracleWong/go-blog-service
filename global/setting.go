package global

import (
	"github.com/MiracleWong/go-blog-service/pkg/logger"
	"github.com/MiracleWong/go-blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)