/*
 * File: \pkg\app\form.go                                                      *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 20:51:32                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 21:18:42                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

func (v *ValidError) Error() string {
	return v.Message
}

type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string

	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)

		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}

	return true, nil
}
