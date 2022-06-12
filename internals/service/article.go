/*
 * File: \internals\service\article.go                                         *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 23:45:37                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/12 , 14:29:36                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package service

import (
	"go_start/blog_service/internals/model"
	"go_start/blog_service/pkg/app"
)

type ArticleGetRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// type ArticleGetByTitleRequest struct {
// 	Title string `form:"title" binding:"required,max=100"`
// 	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
// }

type ArticleCountRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,max=255"`
	Content       string `form:"content" binding:"required"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,max=255"`
	CreatedBy     string `json:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=3,max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255,url"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	Title string `form:"title" binding:"min=3,max=100"`
}

func (svc *Service) GetArticle(param *ArticleGetRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID, "", param.State)
}

// func (svc *Service) GetArticleByTitle(param *ArticleGetByTitleRequest) (*model.Article, error) {
// 	return svc.dao.GetArticle(0,param.Title, param.State)
// }

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CountArticle(param *ArticleCountRequest) (int64, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.State, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
