/*
 * File: \internals\dao\tag.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/05/30 , 21:59:15                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:08:40                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package dao

import (
	"go_start/blog_service/internals/model"
	"go_start/blog_service/pkg/app"
)

func (d *Dao) GetTag(id uint32, state uint8) (*model.Tag, error) {
	tag := model.Tag{
		Model: model.Model{
			ID: id,
		},
		State: state,
	}

	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{State: state}
	pageOffSet := app.GetPageOffSet(page, pageSize)

	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CountTag(state uint8) (int64, error) {
	tags := model.Tag{State: state}

	return tags.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tags := model.Tag{
		Name:  name,
		State: state,
		Model: model.Model{CreatedBy: createdBy},
	}

	return tags.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tags := model.Tag{
		Model: model.Model{
			ID: id,
		},
	}
	values := map[string]any{
		"state":      state,
		"updated_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}

	return tags.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tags := model.Tag{
		Model: model.Model{ID: id},
	}

	return tags.Delete(d.engine)
}

func (d *Dao) GetTagsOfArticle(articleId uint32) ([]*model.Tag, error) {
	tag := model.Tag{}

	return tag.GetTagsOfArticle(d.engine, articleId)
}
