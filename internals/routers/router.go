/*
 * File: \internal\routers\routers.go                                          *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:36:09                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 16:13:36                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package routers

import (
	v1 "go_start/blog_service/internals/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")

	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.List)
		apiV1.GET("/tags", tag.Get)
	}

	return r
}
