/*
 * File: /internals/service/auth.go                                            *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:17:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 10:21:12                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)

	if err != nil {
		return err
	}

	if auth.ID > 0{
		return nil
	}

	return errors.New("auth info does not exists")
}
