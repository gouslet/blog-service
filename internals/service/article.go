/*
 * File: \internals\service\article.go                                         *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 23:45:37                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 04:31:43                             *
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
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"
)

type ArticleGetRequest struct {
	ID uint32 `uri:"id" binding:"required,gte=1"`
}

// type ArticleGetByTitleRequest struct {
// 	Title string `form:"title" binding:"required,max=100"`
// 	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
// }

type ArticleCountRequest struct {
	State uint8 `form:"state,default=2" binding:"oneof=0 1 2"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=2" binding:"oneof=0 1 2"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,max=255"`
	Content       string `form:"content" binding:"required"`
	CoverImageUrl string `json:"cover_image_url" binding:"required,max=255"`
	CreatedBy     string `json:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `uri:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=3,max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255,url"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	UpdatedBy     string `form:"updated_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `uri:"id" binding:"required,gte=1"`
}

// type Article struct {
// 	model.Article
// 	Tags []*model.Tag `json:"tags"`
// }

// type ArticleList struct {
// 	Articles []Article `json:"articles"`
// 	Pager    app.Pager `json:"pager"`
// }

func (svc *Service) GetArticle(param *ArticleGetRequest) (model.ArticleWithTags, error) {
	article, err := svc.dao.GetArticle(param.ID)
	if errors.IsCode(err, errcode.ErrArticleNotFound) {
		return model.ArticleWithTags{}, errors.WrapC(err, errcode.InvalidParams, "")
	}

	// tags, err := svc.dao.GetTagsOfArticle(param.ID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// articleTags := &Article{
	// 	article,
	// 	tags,
	// }

	return article, nil
}

// func (svc *Service) GetArticleByTitle(param *ArticleGetByTitleRequest) (*model.Article, error) {
// 	return svc.dao.GetArticle(0, param.State)
// }

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) (model.ArticleList, error) {
	articleList, err := svc.dao.GetArticleList(param.State, pager)
	if err != nil {
		return model.ArticleList{}, err
	}

	articleList.Pager = pager

	return articleList, nil
}

func (svc *Service) CountArticle(param *ArticleCountRequest) (int64, error) {
	return svc.dao.CountArticle(param.State)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) (model.ArticleWithTags, error) {
	article, err := svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.CreatedBy)

	return article, err
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) (model.ArticleWithTags, error) {
	article, err := svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.UpdatedBy)

	if errors.IsCode(err, errcode.ErrArticleNotFound) {
		return model.ArticleWithTags{}, errors.WrapC(err, errcode.InvalidParams, "")
	}

	return article, err
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) (model.ArticleWithTags, error) {
	article, err := svc.dao.DeleteArticle(param.ID)

	if errors.IsCode(err, errcode.ErrArticleNotFound) {
		return model.ArticleWithTags{}, errors.WrapC(err, errcode.InvalidParams, "")
	}
	return article, err
}
