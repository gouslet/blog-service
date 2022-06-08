/*
 * File: /internals/middleware/limiter.go                                      *
 * Project: blog-service                                                       *
 * Created At: Wednesday, 2022/06/8 , 09:48:23                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 09:51:30                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package middleware

import (
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/errcode"
	"go_start/blog_service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket,ok := l.GetBucket(key);ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)

				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}

