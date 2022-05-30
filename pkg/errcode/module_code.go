/*
 * File: \pkg\errcode\module_code.go                                           *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 23:56:41                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 23:59:59                                *
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
)
