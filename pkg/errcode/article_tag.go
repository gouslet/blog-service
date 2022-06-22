/*
 * File: /pkg/errcode/code.go                                                  *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/06/19 , 15:51:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:46:55                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package errcode

//go:generate codegen -type=int

// blog-service article_tag error
const (
	// ErrArticleTagNotFound - 400: ArticleTag not found.
	ErrArticleTagNotFound int = iota + 110201

	// ErrArticleTagAlreadyExist - 400: ArticleTag already exists.
	ErrArticleTagAlreadyExist
)
