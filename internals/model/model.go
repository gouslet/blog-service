/*
 * File: \internal\model\model.go                                              *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:25:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 15:15:46                                 *
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
	"reflect"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"
)

// Model common model
type Model struct {
	ID         uint32                `gorm:"primary_key" json:"id"`
	CreatedBy  string                `json:"created_by"`
	ModifiedBy string                `json:"modified_by"`
	CreatedAt  uint32                `json:"created_at"`
	ModifiedAt uint32                `json:"modified_at"`
	DeletedAt  uint32                `json:"deleted_at"`
	IsDel      soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
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

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	statement := db.Statement
	context := statement.Context
	for _, filed := range []string{"ModifiedAt", "CreatedAt"} {
		timeField := statement.Schema.LookUpField(filed)
		if !timeField.NotNull {
			reflectValue := statement.ReflectValue
			switch reflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < reflectValue.Len(); i++ {
					if _, isZero := timeField.ValueOf(context, reflectValue.Index(i)); isZero {
						timeField.Set(context, reflectValue.Index(i), time.Now())
					}
				}
			case reflect.Struct:
				if _, isZero := timeField.ValueOf(context, reflectValue); isZero {
					timeField.Set(context, reflectValue, time.Now())
				}
			}
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	statement := db.Statement

	timeField := statement.Schema.LookUpField("gorm:update_column")
	if !timeField.NotNull {
		reflectValue := statement.ReflectValue
		switch reflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < reflectValue.Len(); i++ {
				timeField.Set(statement.Context, reflectValue.Index(i), time.Now())
			}
		case reflect.Struct:
			timeField.Set(statement.Context, reflectValue, time.Now())
		}
	}
}
