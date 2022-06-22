/*
 * File: \pkg\errcode\errcode.go                                               *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 16:17:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 04:49:55                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errcode2

import (
	"fmt"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d existed already, want another", code))
	}

	codes[code] = msg

	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = append([]string{}, details...)

	return &newError
}

// func (e *Error) StatusCode() int {
// 	switch e.code {
// 	case Success.code:
// 		return http.StatusOK
// 	case ServerError.code:
// 		return http.StatusInternalServerError
// 	case InvalidParams.code:
// 		return http.StatusBadRequest
// 	case NotFound.code:
// 		return http.StatusNotFound
// 	case TooManyRequests.code:
// 		return http.StatusTooManyRequests
// 	}

// 	return http.StatusInternalServerError
// }
