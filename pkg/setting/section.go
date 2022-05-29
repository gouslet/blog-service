/*
 * File: \pkg\setting\secontion.go                                             *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:19:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 17:25:32                                *
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
	LogFileName string
	LogFileExt  string
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
