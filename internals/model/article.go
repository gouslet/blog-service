/*
 * File: \internals\model\article.go                                           *
 * Project: blog-service                                                       *
 * Created At: Thursday, 2022/06/2 , 17:58:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/12 , 14:31:08                                *
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

	"gorm.io/gorm"
)

type Article struct {
	*Model

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (*Article, error) {
	db = db.Model(a).Where("id = ? AND title = ? AND is_del = ? AND state = ?", a.ID, 0, a.State)
	if err := db.First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid article_id,not existed")
		}
		return nil, err
	}

	return &a, nil
}

func (a Article) Count(db *gorm.DB) (int64, error) {
	var count int64

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	db = db.Where("state = ?", a.State)

	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) Create(db *gorm.DB) error {
	db = db.Model(a).Where("title = ? AND  is_del = ?", a.Title, 0)

	if err := db.First(&a).Error; err == nil {
		return fmt.Errorf("article: %v has been existed in the system", a.Title)
	}

	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values any) error {
	db = db.Model(a).Where("id = ? AND  is_del = ?", a.ID, 0)
	if err := db.First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("invalid article_id,not existed")
		}
		return err
	}

	return db.Select("*").Updates(values).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ?", a.ID).Delete(&a).Error
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	db = db.Where("state = ?", a.State)

	if err = db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
