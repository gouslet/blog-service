/*
 * File: /pkg/errcode/upload.go                                                *
 * Project: blog-service                                                       *
 * Created At: Wednesday, 2022/06/22 , 07:28:37                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 07:30:00                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package errcode

//go:generate codegen -type=int

// blog-service upload file error
const (
	// ErrUploadFile - 400: Failed uploading files.
	ErrUploadFile int = iota + 120001
)
