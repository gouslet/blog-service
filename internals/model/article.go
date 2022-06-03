/*
 * File: \internals\model\article.go                                           *
 * Project: blog_service                                                       *
 * Created At: Thursday, 2022/06/2 , 17:58:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 09:46:31                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import "gorm.io/gorm"

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
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB) error {
	db = db.Model(&Article{}).Where("id = ?", a.ID)
	return db.Save(&a).Error
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
