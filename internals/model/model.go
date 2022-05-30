/*
 * File: \internal\model\model.go                                              *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:25:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 21:50:14                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import (
	"fmt"
	"go_start/blog_service/global"
	"go_start/blog_service/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Model common model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	if sdb, err := db.DB(); err != nil {
		sdb.SetMaxOpenConns(databaseSetting.MaxIdOpenConns)
	}

	if sdb, err := db.DB(); err != nil {
		sdb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	}

	return db, nil
}
