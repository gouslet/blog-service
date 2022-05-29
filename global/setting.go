/*
 * File: \global\setting.go                                                    *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:26:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 17:36:21                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package global

import "go_start/blog_service/pkg/setting"

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
)
