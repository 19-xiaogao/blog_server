package global

import (
	"github.com/jinzhu/gorm"
	"xiaolong_blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
)
