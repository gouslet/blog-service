/*
 * File: \global\db.go                                                         *
 * Project: blog_service                                                       *
 * Created At: Monday, 2022/05/30 , 17:23:26                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/30 , 17:23:57                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package global

import "gorm.io/gorm"

var (
	DBEngine *gorm.DB
)
