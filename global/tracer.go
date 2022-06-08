/*
 * File: /global/tracer.go                                                     *
 * Project: blog-service                                                       *
 * Created At: Wednesday, 2022/06/8 , 06:50:35                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 06:57:39                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */



package global

import "github.com/opentracing/opentracing-go"

var (
	Tracer opentracing.Tracer
)
