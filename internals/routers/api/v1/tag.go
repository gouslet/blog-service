/*
 * File: \internal\routers\api\v1\tag.go                                       *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:40:25                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 15:32:20                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package v1

import (
	"go_start/blog_service/internals/service"
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/convert"
	"go_start/blog_service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary get a list of tags
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enum(0, 1) default(1)
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} model.Tag "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

// @Summary create a new tag
// @Produce json
// @Param name body string true "tag name" maxlength(100)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.CreateTag(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
	}

	response.ToResponse(gin.H{})
}

// @Summary update a tag
// @Produce json
// @Param id path string true "tag id"
// @Param name body string false "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param modified_by body string true "modifier" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UpdateTag(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
	}

	response.ToResponse(gin.H{})
}

// @Summary delete a tag
// @Produce json
// @Param id path string true "tag id"
// @Success 200 {object} model.Tag "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.DeleteTag(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
	}

	response.ToResponse(gin.H{})
}
