/*
 * File: \internals\dao\tag.go                                                 *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 21:59:15                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/1 , 20:12:42                              *
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
		Name:  name,
		State: state,
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}

	return tags.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tags := model.Tag{
		Model: &model.Model{ID: id},
	}

	return tags.Delete(d.engine)
}
