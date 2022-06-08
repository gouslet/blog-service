/*
 * File: \pkg\setting\secontion.go                                             *
 * Project: blog-service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:19:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 09:11:54                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package setting

import "time"

var sections = make(map[string]any)

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

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
