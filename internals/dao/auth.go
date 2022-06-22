/*
 * File: /internals/dao/jwt.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:15:07                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 07:08:12                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package dao

import (
	"go_start/blog_service/internals/model"
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"
	"gorm.io/gorm"
)

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}

	a, err := auth.Get(d.engine)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Auth{}, errors.WithCode(errcode.ErrAuthNotFound, "auth {key: %q, secret: %q} not found", appKey, appSecret)
	}
	return a, nil
}
