/*
 * File: \internal\routers\api\v1\article.go                                       *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:40:25                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 13:27:36                                 *
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

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary get an article with its title or id
// @Produce json
// @Param id path int true "article id"
// @Param name query string false "article name" maxlength(100)
// @Param state query int false "state" Enum(0, 1) default(1)
// @Success 200 {object} model.Article "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary get a list of articles
// @Produce json
// @Param name query string false "article name" maxlength(100)
// @Param state query int false "state" Enum(0, 1) default(1)
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} model.Article "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}

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
	totalRows, err := svc.CountArticle(&service.CountArticleRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}

	articles, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articles, totalRows)
}

// @Summary create a new article
// @Produce json
// @Param title body string true "Article title" minlength(3) maxlength(100)
// @Param desc body string true "Article description" maxlength(255)
// @Param content body string true "Article content"
// @Param cover_image_url body string true "Article cover image url" maxlength(255)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.CreateArticle(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
	}

	response.ToResponse(gin.H{})
}

// @Summary update an article
// @Produce json
// @Param id path int true "article id"
// @Param title body string false "Article title" minlength(3) maxlength(100)
// @Param desc body string false "Article description" maxlength(255)
// @Param content body string false "Article content"
// @Param cover_image_url body string false "Article cover image url" maxlength(255)
// @Param state body int false "state" Enum(0, 1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UpdateArticle(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
	}

	response.ToResponse(gin.H{})
}

// @Summary delete an article with given title or id
// @Produce json
// @Param id path int true "article id"
// @Param name query string false "article name" maxlength(100)
// @Success 200 {object} model.Article "succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
// @Router       /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, err1 := app.BindAndValid(c, &param)

	if !valid {
		errResp := errcode.InvalidParams.WithDetails(err1.Errors()...)
		response.ToErrorResponse(errResp)
	}

	svc := service.New(c.Request.Context())
	err2 := svc.DeleteArticle(&param)
	if err2 != nil {
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
	}

	response.ToResponse(gin.H{})
}
