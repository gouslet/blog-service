/*
 * File: /internals/model/auth.go                                              *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 09:39:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 10:14:45                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import "gorm.io/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth

	db = db.Where(
		"app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0,
	)

	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return auth,err
	}

	return auth,nil
}
