/*
 * File: /pkg/upload/file.go                                                   *
 * Project: blog-service                                                       *
 * Created At: Monday, 2022/06/6 , 14:41:29                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/6 , 15:22:50                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package upload

import (
	"go_start/blog_service/global"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage = iota + 1

// GetFileName return md5 encrypted name of the file
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	return fileName + ext
}

// GetFileExt return the extension name of the file
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath return the path to save the file
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath check the existence of the dst path
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckPermission the permission of the dst path
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt check if the extension of the file is permitted
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileName(name)
	ext = strings.ToUpper(ext)

	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// CheckMaxSize check if the file size exceed the upper limit
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// CreateSavePath make the directory to save the file
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile save the file
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	out, err := os.Create(dst)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	return err

}
