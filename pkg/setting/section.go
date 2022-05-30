/*
 * File: \pkg\setting\secontion.go                                             *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:19:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/31 , 00:17:08                               *
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
	LogPath         string
	LogFileName     string
	LogFileExt      string
	DefaultPageSize int
	MaxPageSize     int
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

func (s *Setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
