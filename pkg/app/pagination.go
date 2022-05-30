/*
 * File: \pkg\app\pagination.go                                                *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 22:03:17                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/31 , 00:29:09                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package app

import (
	"go_start/blog_service/global"
	"go_start/blog_service/pkg/convert"

	"github.com/gin-gonic/gin"
)

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func GetPageOffSet(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()

	if page < 1 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()

	if pageSize < 1 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}
