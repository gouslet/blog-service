/*
 * File: /pkg/errcode/auth.go                                                  *
 * Project: blog-service                                                       *
 * Created At: Wednesday, 2022/06/22 , 06:58:02                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 07:30:29                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errcode

//go:generate codegen -type=int

// blog-service auth error
const (
	// ErrAuthNotFound - 400: Auth not found.
	ErrAuthNotFound int = iota + 110301

	// ErrAuthAlreadyExist - 400: Auth already exists.
	ErrAuthAlreadyExist
)
