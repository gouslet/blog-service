/*
 * File: \internals\model\article.go                                           *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 17:58:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 16:43:45                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import (
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"

	"gorm.io/gorm"
)

type Article struct {
	Model `gorm:"embedded"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleList struct {
	Articles []ArticleWithTags `json:"articles"`
	Pager    *app.Pager        `json:"pager"`
}

type ArticleWithTags struct {
	Article `gorm:"embedded"`
	Tags    []Tag `json:"tags" gorm:"many2many:blog_article_tag;joinForeignKey:article_id;omitempty;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (awt ArticleWithTags) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (ArticleWithTags, error) {
	article := ArticleWithTags{
		Article: a,
	}

	// if err := db.First(&article).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return ArticleWithTags{}, errors.WithCode(errcode.ErrArticleNotFound,"")
	// 	}
	// 	return ArticleWithTags{}, err
	// }
	err := db.First(&article).Error

	return article, err
}

func (a Article) Count(db *gorm.DB) (int64, error) {
	var count int64

	if a.State != 2 {
		db = db.Where("state = ?", a.State)
	} else {
		db = db.Where("state IN (0,1)")
	}
	// db = db.Where("state = ?", a.State)

	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) Create(db *gorm.DB) (ArticleWithTags, error) {
	article := ArticleWithTags{
		Article: a,
	}

	rows := db.Where(&article).Find(&article).RowsAffected
	if rows != 0 {
		return article, errors.WithCode(errcode.ErrArticleAlreadyExist, "article with title %q already exists", a.Title)
	}
	db.Error = nil
	err := db.Create(&article).Error

	return article, err
}

func (a Article) Update(db *gorm.DB, values any) (ArticleWithTags, error) {
	article := ArticleWithTags{
		Article: a,
	}

	err := db.First(&article).Error
	if err != nil {
		return ArticleWithTags{}, err
	}

	err = db.Model(&article).Updates(values).Error

	return article, err
}

func (a Article) Delete(db *gorm.DB) (ArticleWithTags, error) {
	article := ArticleWithTags{
		Article: a,
	}

	err := db.First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ArticleWithTags{}, err
	}
	err = db.Delete(&article).Error

	return article, err
}

func (a Article) List(db *gorm.DB, pager *app.Pager) (ArticleList, error) {
	var articles []ArticleWithTags
	var err error
	pageOffSet := app.GetPageOffSet(pager.Page, pager.PageSize)

	if pageOffSet >= 0 && pager.PageSize > 0 {
		db = db.Offset(pageOffSet).Limit(pager.PageSize)
	}

	db = db.Preload("Tags", "is_del = ?", 0)

	if a.State == 2 {
		db = db.Where("state IN (0,1)")
	} else {
		db = db.Where("state = ?", a.State)
	}

	if err = db.Where(&ArticleWithTags{}).Find(&articles).Error; err != nil {
		return ArticleList{}, err
	}

	return ArticleList{articles, pager}, nil
}
