package app

import (
	"github.com/MiracleWong/go-blog-service/global"
	"github.com/MiracleWong/go-blog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.DefaultPageSize {
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page <= 0 {
		result = (page - 1) * pageSize
	}
	return result
}