/*
 * File: \internals\dao\tag.go                                                 *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 21:59:15                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/31 , 00:10:11                               *
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
