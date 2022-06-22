/*
 * File: /pkg/errcode/code.go                                                  *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/06/19 , 15:51:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:46:50                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package errcode

//go:generate codegen -type=int

// blog-service article error
const (
	// ErrArticleNotFound - 400: Article not found.
	ErrArticleNotFound int = iota + 110101

	// ErrArticleAlreadyExist - 400: Article already exists.
	ErrArticleAlreadyExist
)
