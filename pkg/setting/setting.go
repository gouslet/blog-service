/*
 * File: \pkg\setting\setting.go                                               *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 17:07:23                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 17:18:47                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yml")
	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
