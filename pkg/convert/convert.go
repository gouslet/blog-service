/*
 * File: \pkg\convert\convert.go                                               *
 * Project: blog_service                                                       *
 * Created At: Tuesday, 2022/05/31 , 00:18:31                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/1 , 20:28:01                              *
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

func (s StrTo) Uint32() (uint32, error) {
	u, err := strconv.ParseUint(s.String(), 10, 64)
	return uint32(u), err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
