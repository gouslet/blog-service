/*
 * File: /internals/middleware/recovery.go                                     *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 14:37:11                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 07:11:40                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package middleware

import (
	"go_start/blog_service/global"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallerFrames().Errorf(ctx, s, err)
				// app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
