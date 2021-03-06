/*
 * File: \internal\routers\api\v1\tag.go                                       *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:40:25                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 08:45:01                             *
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

	"github.com/elchn/errors"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary get a tag with its id
// @Tags tags
// @Produce json
// @Param id path int true "article id"
// @Param state query int false "state" Enum(0, 1) default(1)
// @Success 200 {object} app.SuccessResponse{data=model.Tag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/tas/{id} [get]
func (t Tag) Get(c *gin.Context) {
	param := service.GetTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(errs)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))

		//		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		//		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	tag, err := svc.GetTag(&param)
	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, ""))
		//		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(tag, nil)
}

// @Summary get a list of tags
// @Tags tags
// @Produce json
// @Param state query int false "state" Enum(0, 1, 2) default(2)
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} app.SuccessResponse{data=model.TagList} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(errs)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))

		//		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		//		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		State: param.State,
	})
	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, ""))
		//		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	pager.TotalRows = int(totalRows)

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		//		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponse(tags, nil)
}

// @Summary create a new tag
// @Tags tags
// @Produce json
// @Param name body string true "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.Tag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return
		//		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		//		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.CreateTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
		//		response.ToErrorResponse(errcode.ErrorCreateTagFail)
	}

	response.ToResponse(c, nil)
}

// @Summary update a tag
// @Tags tags
// @Produce json
// @Param id path string true "tag id"
// @Param name body string false "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param updated_by body string true "modifier" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.Tag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)

		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return
		//		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		//		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UpdateTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
		//		response.ToErrorResponse(errcode.ErrorUpdateTagFail.WithDetails(err2.Error()))
	}

	response.ToResponse(c, nil)
}

// @Summary delete a tag
// @Tags tags
// @Produce json
// @Param id path string true "tag id"
// @Success 200 {object} app.SuccessResponse{data=model.Tag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return
		//		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		//		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.DeleteTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
		//		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
	}

	response.ToResponse(c, nil)
}
