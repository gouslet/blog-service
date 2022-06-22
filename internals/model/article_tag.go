/*
 * File: \internals\model\article.go                                           *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 17:58:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:24:20                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import (
	"errors"
	"fmt"
	"go_start/blog_service/pkg/app"

	"gorm.io/gorm"
)

type ArticleTag struct {
	Model

	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

type ArticleTagList struct {
	List  []ArticleTag `json:"article_tags"`
	Pager *app.Pager `json:"pager"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) Get(db *gorm.DB) (*ArticleTag, error) {
	db = db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0)

	if err := db.First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid article_tag_id,not existed")
		}
		return nil, err
	}

	return &a, nil
}

func (a ArticleTag) Count(db *gorm.DB) (int64, error) {
	var count int64

	db = db.Where("is_del = ?", 0)

	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a ArticleTag) Create(db *gorm.DB) error {
	db = db.Model(a).Where("tag_id = ? AND article_id = ? AND is_del = ?", a.TagID, a.ArticleID, 0)
	if err := db.First(&a).Error; err == nil {
		return fmt.Errorf("article_tag: %v has been existed in the system", a)
	}
	db.Error = nil

	return db.Create(&a).Error
}

func (a ArticleTag) Update(db *gorm.DB, values any) error {
	db = db.Model(a).Where("id = ? AND  is_del = ?", a.ID, 0)
	if err := db.First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("invalid article_tag_id,not existed")
		}
		return err
	}

	return db.Select("*").Updates(values).Error
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", a.ID).Delete(&a).Error
}

func (a ArticleTag) List(db *gorm.DB, pageOffset, pageSize int) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	db = db.Where("is_del = ?", 0)

	if err = db.Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}
