/*
 * File: /internals/dao/jwt.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 10:15:07                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 10:16:46                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package dao

import "go_start/blog_service/internals/model"

func (d *Dao)GetAuth(appKey,appSecret string)(model.Auth,error) {
	auth := model.Auth{
		AppKey: appKey,
		AppSecret: appSecret,
	}

	return auth.Get(d.engine)
}

