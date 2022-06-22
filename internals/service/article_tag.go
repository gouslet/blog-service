/*
 * File: \internals\service\article.go                                         *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 23:45:37                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:08:40                               *
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

type ArticleTagGetRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// type ArticleTagGetByTitleRequest struct {
// 	Title string `form:"title" binding:"required,max=100"`
// 	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
// }

type ArticleTagCountRequest struct {
}

type ArticleTagListRequest struct {
}

type CreateArticleTagRequest struct {
	TagID     uint32 `form:"tag_id" binding:"required,gte=1"`
	ArticleID uint32 `form:"tag_id" binding:"required,gte=1"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleTagRequest struct {
	ID        uint32 `form:"id" binding:"required,gte=1"`
	TagID     uint32 `form:"tag_id" binding:"gte=1"`
	ArticleID uint32 `form:"tag_id" binding:"gte=1"`
	UpdatedBy string `form:"updated_by" binding:"required,min=3,max=100"`
}

type DeleteArticleTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticleTag(param *ArticleTagGetRequest) (*model.ArticleTag, error) {
	return svc.dao.GetArticleTag(param.ID)
}

// func (svc *Service) GetArticleTagByTitle(param *ArticleTagGetByTitleRequest) (*model.ArticleTag, error) {
// 	return svc.dao.GetArticleTag(0, param.State)
// }

func (svc *Service) GetArticleTagList(param *ArticleTagListRequest, pager *app.Pager) ([]*model.ArticleTag, error) {
	return svc.dao.GetArticleTagList(pager.Page, pager.PageSize)
}

func (svc *Service) CountArticleTag(param *ArticleTagCountRequest) (int64, error) {
	return svc.dao.CountArticleTag()
}

func (svc *Service) CreateArticleTag(param *CreateArticleTagRequest) error {
	return svc.dao.CreateArticleTag(param.ArticleID, param.TagID, param.CreatedBy)
}

func (svc *Service) UpdateArticleTag(param *UpdateArticleTagRequest) error {
	return svc.dao.UpdateArticleTag(param.ID, param.ArticleID, param.TagID, param.UpdatedBy)
}

func (svc *Service) DeleteArticleTag(param *DeleteArticleTagRequest) error {
	return svc.dao.DeleteArticleTag(param.ID)
}
