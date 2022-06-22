/*
 * File: /internals/service/auth.go                                            *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:17:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/22 , 07:10:30                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package service

import (
	"go_start/blog_service/pkg/errcode"

	"github.com/elchn/errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)

	if errors.IsCode(err, errcode.ErrAuthNotFound) {
		return errors.WrapC(err, errcode.UnauthorizedAuthNotExist, "auth info does not exists")
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.WithCode(errcode.UnauthorizedAuthNotExist, "auth info does not exists")

}
