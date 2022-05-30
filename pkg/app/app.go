/*
 * File: \pkg\app\app.go                                                       *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 17:37:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 17:46:31                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package app

import (
	"go_start/blog_service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}

	details := err.Details()

	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}