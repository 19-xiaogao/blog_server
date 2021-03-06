package global

import (
	"github.com/jinzhu/gorm"
	"xiaolong_blog/pkg/Logger"
	"xiaolong_blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
