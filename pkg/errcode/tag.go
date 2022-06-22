/*
 * File: /pkg/errcode/code.go                                                  *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/06/19 , 15:51:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:42:49                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package errcode

//go:generate codegen -type=int

// blog-service tag error
const (
	// ErrTagNotFound - 400: Tag not found.
	ErrTagNotFound int = iota + 110001

	// ErrTagAlreadyExist - 400: Tag already exists.
	ErrTagAlreadyExist
)
