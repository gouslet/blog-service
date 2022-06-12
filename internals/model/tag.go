/*
 * File: \internal\model\model.go                                              *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:25:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/06/11 , 15:56:42                              *
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

// Tag Model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
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

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	db = db.Model(t).Where("name = ? AND  is_del = ?", t.Name, 0)

	if err := db.First(&t).Error; err == nil {
		return fmt.Errorf("tag: %v has been existed in the system",t.Name)
	}

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
