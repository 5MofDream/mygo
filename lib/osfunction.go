package lib

import (
	"os"
	"io/ioutil"
	"fmt"
)

// 检测目录
func PathExists(path string) (bool) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	} else {
		return false
	}
}

// 检测文件
func FileExists(filename string) (bool) {
	info, err := os.Stat(filename)
	if err == nil && info.Mode().IsRegular() {
		return true
	} else {
		return false
	}
}

func GetAllFileByPath(path string) ([]string, error) {
	dirList, err := ioutil.ReadDir(path)
	if (err != nil) {
		return nil, err
	}
	var fileList []string
	for _, v := range dirList {
		switch mode := v.Mode(); {
		case mode.IsDir():
			childFileList, err := GetAllFileByPath(path + "/" + v.Name())
			if (err != nil) {
				panic(fmt.Sprintf("get child dir file error:%v", err))
			}
			fileList = append(fileList, childFileList...)
		case mode.IsRegular():
			fileList = append(fileList, path+"/"+v.Name())
		}
	}
	return fileList, err
}
