/*
 * File: \pkg\errcode.go                                                       *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 16:16:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 16:54:04                                *
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
)
