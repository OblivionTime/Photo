/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-06-14 18:25:19
 * @LastEditors: solid
 * @LastEditTime: 2022-07-11 10:44:01
 */
package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func String2Number(str string) uint {
	if str == "" {
		return 0
	}
	num, _ := strconv.Atoi(str)
	return uint(num)
}
func GetImgFiles(pathSrc string, imgType string) []string {
	files, err := ioutil.ReadDir(pathSrc)
	if err != nil {
		log.Fatal(err)
	}
	var imgFileList []string
	flag := true
	if imgType == "" {
		flag = false
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := path.Ext(file.Name())
		if !flag {
			if ext == ".webp" || ext == ".webp2" || ext == ".bmp" || ext == ".jpeg" || ext == ".jpg" || ext == ".png" || ext == ".tiff" {
				imgFileList = append(imgFileList, file.Name())
			}
		} else {
			if ext == imgType {
				imgFileList = append(imgFileList, file.Name())
			}
		}
	}
	return imgFileList
}
