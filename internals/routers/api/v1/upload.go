/*
 * File: /internals/routers/upload.go                                          *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/06/6 , 15:28:44                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/6 , 15:37:16                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package api

import "github.com/gin-gonic/gin"

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)
