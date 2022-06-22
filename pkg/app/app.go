/*
 * File: \pkg\app\app.go                                                       *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/05/30 , 17:37:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 08:41:57                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package app

import (
	"net/http"

	"github.com/elchn/errors"
	"github.com/gin-gonic/gin"
)

type Error struct {
	code    int     `json:"code"`
	message string  `json:"message"`
	details []Error `json:"details,omitempty"`
}

// SuccessResponse defines the return messages when no error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type SuccessResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Data any `json:"data,omitempty"`
}

// ErrorResponse defines the return messages when a error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type ErrorResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	Errors errors.MyError `json:"errors,omitempty"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
// func WriteResponse(c *gin.Context, err error, data interface{}) {
// 	if err != nil {
// 		coder := errors.ParseCoder(err)

// 		c.JSON(coder.HTTPStatus(), ErrResponse{
// 			Code:      coder.Code(),
// 			Message:   coder.String(),
// 			Reference: coder.Reference(),
// 			// Details:   errors.Cause(err),
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, data)
// }

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func (r *Response) ToResponse(data any, err error) {
	if err != nil {

		coder := errors.ParseCoder(err)
		r.Ctx.JSON(coder.HTTPStatus(), ErrorResponse{
			Code:    coder.HTTPStatus(),
			Message: coder.String(),
			Errors:  errors.ToMyError(errors.Unwrap(err)),
		})
		return
	}

	r.Ctx.JSON(http.StatusOK, SuccessResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

// func (r *Response) ToErrorResponse(err *errcode.Error) {
// 	response := gin.H{
// 		"code": err.Code(),
// 		"msg":  err.Msg(),
// 	}

// 	details := err.Details()

// 	if len(details) > 0 {
// 		response["details"] = details
// 	}

// 	r.Ctx.JSON(err.StatusCode(), response)
// }

// func (r *Response) ToResponseList(list any, totalRows int64) {
// 	r.Ctx.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"list": list,
// 			"pager": Pager{
// 				Page:      GetPage(r.Ctx),
// 				PageSize:  GetPageSize(r.Ctx),
// 				TotalRows: int(totalRows),
// 			},
// 		},
// 	)
// }
