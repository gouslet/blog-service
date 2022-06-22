/*
 * File: \internal\model\model.go                                              *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:25:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 16:52:05                               *
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

// Tag Model
type Tag struct {
	Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagList struct {
	List  []Tag      `json:"tags"`
	Pager *app.Pager `json:"pager"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	db = db.Where("id = ? AND state = ?", t.ID, t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64

	if t.State != 2 {
		db = db.Where("state = ?", t.State)
	} else {
		db = db.Where("state IN (0,1)")
	}

	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.State == 2 {
		db = db.Where("state IN (0,1)")
	} else {
		db = db.Where("state = ?", t.State)
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	db = db.Model(t).Where("name = ? AND  is_del = ?", t.Name, 0)

	if err := db.First(&t).Error; err == nil {
		return fmt.Errorf("tag: %v has been existed in the system", t.Name)
	}

	db.Error = nil

	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values any) error {
	db = db.Model(t).Where("id = ? AND  is_del = ?", t.ID, 0)
	if err := db.First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("invalid tag_id,not existed")
		}
		return err
	}

	return db.Select("*").Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del", t.ID, 0).Delete(&t).Error
}

func (t Tag) GetTagsOfArticle(db *gorm.DB, articleId uint32) ([]*Tag, error) {
	var tags []*Tag

	db.Model(&t).Select("blog_tag.id,blog_tag.name,blog_tag.state").Joins("LEFT JOIN blog_article_tag ON blog_tag.id = blog_article_tag.tag_id").Joins("LEFT JOIN blog_article ON blog_article_tag.article_id = blog_article.id").Scan(tags)
	// SELECT blog_tag.id,blog_tag.name,blog_tag.state FROM blog_tag LEFT JOIN blog_article_tag ON blog_tag.id = blog_article_tag.tag_id LEFT JOIN blog_article ON blog_article_tag.article_id = blog_article.id

	return tags, nil
}
