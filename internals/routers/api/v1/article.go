/*
 * File: \internal\routers\api\v1\article.go                                       *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:40:25                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 08:43:46                             *
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

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary get an article with its id
// @Tags article
// @Produce json
// @Param id path int true "article id"
// @Param state query int false "state" Enum(0, 1) default(1)
// @Success 200 {object} app.SuccessResponse{data=model.ArticleWithTags} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {
	param := service.ArticleGetRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	// errs := c.BindUri(&param)
	// if errs != nil {
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		x := errors.WrapC(
			errs,
			errcode.InvalidParams,
			errs.Error(),
		)
		response.ToResponse(nil, x)
		// errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	article, err := svc.GetArticle(&param)
	response.ToResponse(article, err)
}

// // @Summary get an article with its title
// // @Tags article
// // @Produce json
// // @Param title path string true "article title" maxlength(100)
// // @Param state query int false "state" Enum(0, 1) default(1)
// // @Success 200 {object} model.Article "succeeded"
// // @Failure 400 {object} errcode.Error "request errors"
// // @Failure 500 {object} errcode.Error "internal errors"
// // @Router       /api/v1/articles/{title} [get]
// func (t Article) GetByTitle(c *gin.Context) {
// 	param := service.ArticleGetByTitleRequest{}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)

// 	if !valid {
// 		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
// 		response.ToErrorResponse(errRsp)
// 		return
// 	}

// 	svc := service.New(c.Request.Context())

// 	article, err := svc.GetArticleByTitle(&param)
// 	if err != nil {
// 		response.ToErrorResponse(errcode.ErrorGetArticleFail)
// 		return
// 	}

// 	response.ToResponse(article)
// }

// @Summary get a list of articles
// @Tags article
// @Produce json
// @Param state query int false "state" Enum(0, 1, 2) default(2)
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} app.SuccessResponse{data=model.ArticleList} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
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

	totalRows, err := svc.CountArticle(&service.ArticleCountRequest{State: param.State})
	if err != nil {
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))
		// response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}
	pager.TotalRows = int(totalRows)

	articles, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		fmt.Println(err)
		response.ToResponse(nil, errors.WrapC(errs, errcode.InvalidParams, ""))
		return
		// response.ToResponse(nil, err)
		// response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		// return
	}

	response.ToResponse(articles, err)
}

// @Summary create a new article
// @Tags article
// @Produce json
// @Param title body string true "Article title" minlength(3) maxlength(100)
// @Param desc body string true "Article description" maxlength(255)
// @Param content body string true "Article content"
// @Param cover_image_url body string true "Article cover image url" maxlength(255)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.ArticleWithTags} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WrapC(err1, errcode.InvalidParams, ""))
		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	article, err2 := svc.CreateArticle(&param)

	if errors.IsCode(err2, errcode.ErrArticleAlreadyExist) {
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
	}

	response.ToResponse(article, nil)
}

// @Summary update an article
// @Tags article
// @Produce json
// @Param id path int true "article id"
// @Param title body string false "Article title" minlength(3) maxlength(100)
// @Param desc body string false "Article description" maxlength(255)
// @Param content body string false "Article content"
// @Param cover_image_url body string false "Article cover image url" maxlength(255)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param updated_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} app.SuccessResponse{data=model.ArticleWithTags} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		fmt.Println(err1)
		response.ToResponse(nil, errors.WithCode(
			errcode.InvalidParams,
			err1.Error(),
		))
		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	article, err2 := svc.UpdateArticle(&param)
	if err2 != nil {
		// response.ToResponse(nil,err2)
		fmt.Println(err2)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
		// response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
	}

	response.ToResponse(article, err2)
}

// @Summary delete an article with given title or id
// @Tags article
// @Produce json
// @Param id path int true "article id"
// @Success 200 {object} app.SuccessResponse{data=model.ArticleWithTags} "succeeded"
// @Failure 400 {object} app.ErrorResponse "request errors"
// @Failure 500 {object} app.ErrorResponse "internal errors"
// @Router       /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		response.ToResponse(nil, errors.WithCode(
			errcode.InvalidParams,
			err1.Error(),
		))
		// errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		// response.ToErrorResponse(errResp)
		fmt.Println(err1)
		return
	}

	svc := service.New(c.Request.Context())
	article, err2 := svc.DeleteArticle(&param)
	if err2 != nil {
		fmt.Println(err2)
		// response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		response.ToResponse(nil, errors.WrapC(err2, errcode.InvalidParams, ""))
		return
	}

	response.ToResponse(article, err2)
}
