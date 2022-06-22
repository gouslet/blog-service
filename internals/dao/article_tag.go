/*
 * File: /internals/dao/article_tag.go                                         *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/06/13 , 11:55:51                                   *
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

func (d *Dao) GetArticleTag(id uint32) (*model.ArticleTag, error) {
	articleTag := model.ArticleTag{
		Model: model.Model{
			ID: id,
		},
	}

	return articleTag.Get(d.engine)
}

func (d *Dao) GetArticleTagList(page, pageSize int) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	pageOffSet := app.GetPageOffSet(page, pageSize)

	return articleTag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CountArticleTag() (int64, error) {
	articleTags := model.ArticleTag{}

	return articleTags.Count(d.engine)
}

func (d *Dao) CreateArticleTag(article_id, tag_id uint32, createdBy string) error {
	articleTags := model.ArticleTag{
		Model: model.Model{
			CreatedBy: createdBy,
		},
		ArticleID: article_id,
		TagID:     tag_id,
	}

	return articleTags.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(id, article_id, tag_id uint32, modifiedBy string) error {
	articleTags := model.ArticleTag{
		Model: model.Model{
			ID: id,
		},
	}

	values := map[string]any{
		"updated_by": modifiedBy,
	}

	if article_id != 0 {
		values["article_id"] = article_id
	}

	if tag_id != 0 {
		values["tag_id"] = tag_id
	}

	return articleTags.Update(d.engine, values)
}

func (d *Dao) DeleteArticleTag(id uint32) error {
	articleTags := model.ArticleTag{
		Model: model.Model{ID: id},
	}

	return articleTags.Delete(d.engine)
}
