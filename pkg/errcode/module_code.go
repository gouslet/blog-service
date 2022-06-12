/*
 * File: \pkg\errcode\module_code.go                                           *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/05/30 , 23:56:41                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/12 , 09:59:00                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "failed getting tags list")
	ErrorCreateTagFail  = NewError(20010002, "failed creating a tag")
	ErrorUpdateTagFail  = NewError(20010003, "failed updating a tag")
	ErrorDeleteTagFail  = NewError(20010004, "failed deleting a tag")
	ErrorCountTagFail   = NewError(20010005, "failed counting tags")

	ErrorGetArticleListFail = NewError(20010006, "failed getting articles list")
	ErrorCreateArticleFail  = NewError(20010007, "failed creating an article")
	ErrorUpdateArticleFail  = NewError(20010008, "failed updating an article")
	ErrorDeleteArticleFail  = NewError(20010009, "failed deleting an article")
	ErrorCountArticleFail   = NewError(20010010, "failed counting articles")
	ErrorGetArticleFail = NewError(20010011, "failed getting an article")

	ErrorUploadFileFailed = NewError(2003001,"failed uploading a file")
)
