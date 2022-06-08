/*
 * File: /internals/middleware/access_log.go                                   *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 12:33:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 07:27:09                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package middleware

import (
	"bytes"
	"go_start/blog_service/global"
	"go_start/blog_service/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = bodyWriter

		beginTime := time.Now()
		ctx.Next()
		endTime := time.Now()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		format := "access log: method: %s, status code: %d, begin time: %d, end time %d"
		global.Logger.WithFields(fields).Infof(
			format,
			ctx.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
