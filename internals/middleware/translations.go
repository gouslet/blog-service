/*
 * File: \internals\middleware\translations.go                                 *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 20:58:10                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 21:12:43                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/en"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				zh_translations.RegisterDefaultTranslations(v, trans)
			case "en":
				en_translations.RegisterDefaultTranslations(v, trans)
			default:
				zh_translations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
