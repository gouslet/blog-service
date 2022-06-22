/*
 * File: \internals\dao\article.go                                             *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 23:46:48                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 16:32:46                               *
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
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"

	"gorm.io/gorm"
)

func (d *Dao) GetArticle(id uint32) (model.ArticleWithTags, error) {
	article := model.Article{
		Model: model.Model{
			ID: id,
		},
	}
	awt, err := article.Get(d.engine)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ArticleWithTags{}, errors.WithCode(errcode.ErrArticleNotFound, "article with id %v does not exist", id)
	}

	return awt, nil
}

func (d *Dao) GetArticleList(state uint8, pager *app.Pager) (model.ArticleList, error) {
	article := model.Article{State: state}

	return article.List(d.engine, pager)
}

func (d *Dao) CountArticle(state uint8) (int64, error) {
	articles := model.Article{State: state}

	return articles.Count(d.engine)
}

func (d *Dao) CreateArticle(title, desc, content, coverImageUrl, createdBy string) (model.ArticleWithTags, error) {
	articles := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		Model: model.Model{
			CreatedBy: createdBy,
		},
	}

	return articles.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, desc, content, cover_image_url string, state uint8, modifiedBy string) (model.ArticleWithTags, error) {
	articles := model.Article{
		Model: model.Model{
			ID: id,
		},
	}
	values := map[string]any{
		"state":      state,
		"updated_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}
	if cover_image_url != "" {
		values["cover_image_url"] = cover_image_url
	}
	article, err := articles.Update(d.engine, values)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return article, errors.WithCode(errcode.ErrArticleNotFound, "article with id %v does not exist", id)
	}

	return article, nil
}

func (d *Dao) DeleteArticle(id uint32) (model.ArticleWithTags, error) {
	articles := model.Article{
		Model: model.Model{ID: id},
	}
	article, err := articles.Delete(d.engine)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return article, errors.WithCode(errcode.ErrArticleNotFound, "article with id %v does not exist", id)
	}

	return article, nil
}
