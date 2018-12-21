package upload

import (
	"meifu/conf"
	"path"
	"strings"
	"meifu/util"
	"mime/multipart"
	"io/ioutil"
	"os"
	"fmt"
)

// 	获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return conf.APIURL + GetRootPath() + GetImagePath() + name
}

//	获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//	获取图片路径
func GetImagePath() string {
	return conf.IMAGEPATH
}

//	获取根
func GetRootPath() string {
	return conf.RUNTIMEROOTPATH
}


//	获取图片完整路径
func GetImageFullPath() string {
	return conf.RUNTIMEROOTPATH + GetImagePath()
}

//	检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := path.Ext(fileName) 	// path.Ext获取文件后缀
	for _, allowExt := range conf.IMAGEALLOWEXTS {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {  //  toupper  返回将所有字母都转为对应的大写版本的拷贝。
			return true
		}
	}

	return false
}

//	获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

//	检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return size <= conf.IMAGEMAXSIZE
}

//  检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

//	新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//	如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//	检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

//	检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
