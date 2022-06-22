/*
 * File: /internals/model/model.go                                             *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/06/12 , 15:52:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:09:00                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

/*
 * File: /internals/model/model.go                                             *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/06/12 , 15:52:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:08:40                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

/*
 * File: \internal\model\model.go                                              *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:25:51                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:07:59                               *
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
	ID        uint32                `json:"id" gorm:"primary_key"`
	CreatedBy string                `json:"created_by"`
	UpdatedBy string                `json:"updated_by"`
	CreatedAt uint32                `json:"created_at"`
	UpdatedAt uint32                `json:"updated_at"`
	IsDel     soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_name"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	newLogger := logger.Default
	if global.ServerSetting.RunMode == "debug" {
		newLogger = newLogger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
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
	for _, filed := range []string{"UpdatedAt", "CreatedAt"} {
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
