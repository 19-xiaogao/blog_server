package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xiaolong_blog/global"
)

func GetPage(c *gin.Context) int {

	page, ok := strconv.Atoi(string(c.Query("page")))
	if ok != nil {
		return 0
	}
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, ok := strconv.Atoi(string(c.Query("pageSize")))
	if ok != nil || pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.DefaultPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
