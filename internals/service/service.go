/*
 * File: \internals\service\service.go                                         *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 23:43:38                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 23:45:40                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package service

import (
	"context"
	"go_start/blog_service/global"
	"go_start/blog_service/internals/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)

	return svc
}
