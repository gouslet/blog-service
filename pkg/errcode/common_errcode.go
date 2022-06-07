/*
 * File: \pkg\errcode.go                                                       *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 16:16:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 10:33:58                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errcode

var (
	Success         = NewError(0, "Succeeded")
	ServerError     = NewError(10000000, "Internal error")
	InvalidParams   = NewError(10000001, "Parameters error")
	NotFound        = NewError(10000002, "Not found")
	TooManyRequests = NewError(10000003, "Too many requests")

	UnauthorizedAuthNotExist  = NewError(10000004, "Authorization failed: can't find AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000005, "Authorization failed: token error")
	UnauthorizedTokenTimeout  = NewError(10000006, "Authorization failed: token timeout")
	UnauthorizedTokenGenerate = NewError(10000007, "Authorization failed: failed generating a token")
)
