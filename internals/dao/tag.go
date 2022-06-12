/*
 * File: \internals\dao\tag.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/05/30 , 21:59:15                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/06/11 , 16:00:56                              *
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

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffSet := app.GetPageOffSet(page, pageSize)

	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tags := model.Tag{Name: name, State: state}

	return tags.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tags := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tags.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tags := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]any{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}

	return tags.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tags := model.Tag{
		Model: &model.Model{ID: id},
	}

	return tags.Delete(d.engine)
}
