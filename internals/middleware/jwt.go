/*
 * File: /internals/middleware/jwt.go                                          *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:01:47                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 10:14:34                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package middleware

import (
	"errors"
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/errcode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exist := ctx.GetQuery("token"); exist {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)

			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					ecode = errcode.UnauthorizedTokenTimeout
				} else {
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(ctx)
			response.ToErrorResponse(ecode)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
