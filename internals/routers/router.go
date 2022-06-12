/*
 * File: \internal\routers\routers.go                                          *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:36:09                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/12 , 10:17:12                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package routers

import (
	_ "go_start/blog_service/docs"
	"go_start/blog_service/global"
	"go_start/blog_service/internals/middleware"
	api "go_start/blog_service/internals/routers/api/v1"
	"go_start/blog_service/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	tag := api.NewTag()
	article := api.NewArticle()
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())

	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.GetByID)
		apiV1.GET("/articles/:title", article.GetByTitle)
		apiV1.GET("/articles", article.List)

	}

	upload := api.NewUpload()

	r.POST("/upload/file", upload.UploadFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", api.GetAuth)
	return r
}
