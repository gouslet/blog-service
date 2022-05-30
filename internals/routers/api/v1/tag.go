/*
 * File: \internal\routers\api\v1\tag.go                                       *
 * Project: blog_service                                                       *
 * Created At: Sunday, 2022/05/29 , 00:40:25                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 20:15:10                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package v1

import (
	"go_start/blog_service/pkg/app"
	"go_start/blog_service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary get multiple tags
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enum(0, 1) default(1)
// @Param page query int false "page index"
// @Param page_size query int false "size per page"
// @Success 200 {object} model.Tag "Succeeded"
// @Failure 400 {object} errcode.Error "request errors"
// @Failure 500 {object} errcode.Error "internal errors"
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
