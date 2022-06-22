/*
 * File: /internals/routers/upload.go                                          *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/06/6 , 15:28:44                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 09:37:38                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package api

import (
	"fmt"
	"go_start/blog_service/internals/service"
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/convert"
	"go_start/blog_service/pkg/errcode"
	"go_start/blog_service/pkg/upload"

	"github.com/elchn/errors"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// @Summary upload a file and get access url back
// @Tags upload
// @Param type body int true "file type" Enum(1,2,3) default(1)
// @Param file body string true "file"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} app.SuccessResponse{data=gin.H} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)

	file, fileHeader, err := c.Request.FormFile("file")
	fType := c.PostForm("type")
	fileType := convert.StrTo(fType).MustInt()

	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, err.Error()))
		return

		// errRsp := errcode.InvalidParams.WithDetails(err.Error())
		// response.ToErrorResponse(errRsp)
	}

	if fileHeader == nil || fileType <= 0 {
		// response.ToErrorResponse(errcode.InvalidParams)
		fmt.Println(err)
		x := errors.WithCode(errcode.ErrUploadFile, "file type of %v is not supported", fType)
		response.ToResponse(nil, errors.WrapC(x, errcode.InvalidParams, ""))
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		// errRsp := errcode.ErrorUploadFileFailed.WithDetails(err.Error())
		// response.ToErrorResponse(errRsp)
		fmt.Println(err)
		x := errors.WithCode(errcode.ErrUploadFile, err.Error())
		response.ToResponse(nil, errors.WrapC(x, errcode.InvalidParams, "file type of %s is not supported", fType))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	}, nil)

}
