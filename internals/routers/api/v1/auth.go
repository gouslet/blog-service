/*
 * File: /internals/routers/api/v1/auth.go                                     *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:22:04                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 07:20:37                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package api

import (
	"fmt"
	"go_start/blog_service/internals/service"
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}

	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(errs)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))
		return
		// errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)

		// response.ToErrorResponse(errRsp)
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)

	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))
		return
		
		// response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)

	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))
		return
		// response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	}
	
	response.ToResponse(gin.H{
		"token": token,
	}, nil)

}
