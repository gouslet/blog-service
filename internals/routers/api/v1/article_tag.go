/*
 * File: /internals/routers/api/v1/articletag_tag.go                              *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/06/13 , 12:16:17                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 08:45:43                             *
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

type ArticleTag struct{}

func NewArticleTag() ArticleTag {
	return ArticleTag{}
}

// @Summary get an articletag with its id
// @Tags articletag
// @Produce json
// @Param id path int true "articletag id"
// @Success 200 {object} app.SuccessResponse{data=model.ArticleTag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articletags/{id} [get]
func (t ArticleTag) Get(c *gin.Context) {
	param := service.ArticleTagGetRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(errs)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))

		// errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	articletag, err := svc.GetArticleTag(&param)
	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorGetArticleTagFail)
		return
	}

	response.ToResponse(articletag, nil)
}

// @Summary get a list of articletags
// @Tags articletag
// @Produce json
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} app.SuccessResponse{data=model.ArticleTagList} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articletags [get]
func (a ArticleTag) List(c *gin.Context) {
	param := service.ArticleTagListRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(errs)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))

		// errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountArticleTag(&service.ArticleTagCountRequest{})
	if err != nil {
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorCountArticleTagFail)
		return
	}
	pager.TotalRows = int(totalRows)

	articletags, err := svc.GetArticleTagList(&param, &pager)
	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(err, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorGetArticleTagListFail)
		return
	}

	response.ToResponse(articletags, nil)
}

// @Summary create a new articletag
// @Tags articletag
// @Produce json
// @Param article_id body int true "article id"
// @Param tag_id body int true "tag id"
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.ArticleTag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articletags [post]
func (t ArticleTag) Create(c *gin.Context) {
	param := service.CreateArticleTagRequest{}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return

		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.CreateArticleTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorCreateArticleTagFail)
		return
	}

	response.ToResponse(c, nil)
}

// @Summary update an articletag
// @Tags articletag
// @Produce json
// @Param id path int true "articletag id"
// @Param article_id body int true "article id"
// @Param tag_id body int true "tag id"
// @Param updated_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.ArticleTag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articletags/{id} [put]
func (t ArticleTag) Update(c *gin.Context) {
	param := service.UpdateArticleTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return

		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UpdateArticleTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorUpdateArticleTagFail)
		return
	}

	response.ToResponse(c, nil)
}

// @Summary delete an articletag with given title or id
// @Tags articletag
// @Produce json
// @Param id path int true "articletag id"
// @Success 200 {object} app.SuccessResponse{data=model.ArticleTag} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articletags/{id} [delete]
func (t ArticleTag) Delete(c *gin.Context) {
	param := service.DeleteArticleTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		return

		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.DeleteArticleTag(&param)
	if err2 != nil {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorUpdateArticleTagFail)
		return
	}

	response.ToResponse(c, nil)
}
