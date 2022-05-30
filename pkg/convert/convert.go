/*
 * File: \pkg\convert\convert.go                                               *
 * Project: blog_service                                                       *
 * Created At: Tuesday, 2022/05/31 , 00:18:31                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/31 , 00:23:31                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) Uint64() (uint64, error) {
	return strconv.ParseUint(s.String(), 10, 64)
}

func (s StrTo) MustUint64() uint64 {
	v, _ := s.Uint64()
	return v
}
