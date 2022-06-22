/*
 * File: \internals\service\tag.go                                             *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/05/30 , 20:44:06                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 16:49:03                               *
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

type GetTagRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CountTagRequest struct {
	State uint8 `form:"state,default=2" binding:"oneof=0 1 2"`
}

type TagListRequest struct {
	State uint8 `form:"state,default=2" binding:"oneof=0 1 2"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"max=100"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID        uint32 `form:"id" binding:"required,gte=1"`
	Name      string `json:"name" binding:"min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	UpdatedBy string `json:"updated_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return svc.dao.GetTag(param.ID, param.State)
}

type TagList struct {
	Tags  []*model.Tag
	Pager app.Pager
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) (TagList, error) {
	tags, _ := svc.dao.GetTagList(param.State, pager.Page, pager.PageSize)
	return TagList{tags, *pager}, nil
}

func (svc *Service) CountTag(param *CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.State)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.UpdatedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
