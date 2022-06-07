/*
 * File: \pkg\setting\secontion.go                                             *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:19:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 09:52:11                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package setting

import "time"

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	LogPath              string
	LogFileName          string
	LogFileExt           string
	DefaultPageSize      int
	MaxPageSize          int
	UploadSavePath       string   // 上传文件的保存位置
	UploadServerUrl      string   // 上传文件后用于展示的文件服务地址
	UploadImageMaxSize   int      // 上传文件所允许的最大空间大小，单位：MB
	UploadImageAllowExts []string // 上传文件所允许的文件后缀
}

type DatabaseSettings struct {
	DBType         string
	Username       string
	Password       string
	Host           string
	DBName         string
	TablePrefix    string
	Charset        string
	ParseTime      bool
	MaxIdleConns   int
	MaxIdOpenConns int
}

type JWTSettings struct {
	AppKey    string
	AppSecret string
	Issuer    string
	Expire    time.Duration
}

func (s *Setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
