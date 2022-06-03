/*
 * File: \internals\dao\article.go                                             *
 * Project: blog_service                                                       *
 * Created At: Thursday, 2022/06/2 , 23:46:48                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 09:48:38                                 *
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

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffSet := app.GetPageOffSet(page, pageSize)

	return article.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CountArticle(title string, state uint8) (int64, error) {
	articles := model.Article{Title: title, State: state}

	return articles.Count(d.engine)
}

func (d *Dao) CreateArticle(title string, state uint8, createdBy string) error {
	articles := model.Article{
		Title: title,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return articles.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, state uint8, modifiedBy string) error {
	articles := model.Article{
		Title: title,
		State: state,
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}

	return articles.Update(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	articles := model.Article{
		Model: &model.Model{ID: id},
	}

	return articles.Delete(d.engine)
}
