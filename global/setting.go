/*
 * File: \global\setting.go                                                    *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:26:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 14:25:07                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package global

import (
	"go_start/blog_service/pkg/logger"
	"go_start/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings
	Logger          *logger.Logger
)
