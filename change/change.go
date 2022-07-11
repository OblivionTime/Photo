/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-06-14 18:15:44
 * @LastEditors: solid
 * @LastEditTime: 2022-07-11 18:34:58
 */
package change

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"photo/logger"
	"photo/utils"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// 图片尺寸修改
func CutImageResource(buf []byte, height, width uint, layout string) ([]byte, error) {
	decodeBuf, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		fmt.Println("当前图片不合法")
		return nil, err
	}
	// 修改图片的大小
	set := resize.Resize(width, height, decodeBuf, resize.Lanczos3)
	NewBuf := bytes.Buffer{}
	switch layout {
	case "png":
		err = png.Encode(&NewBuf, set)
	case "jpeg", "jpg":
		err = jpeg.Encode(&NewBuf, set, nil)
	case "webp", "webp2":
		err = webp.Encode(&NewBuf, set, nil)
	case "bmp":
		err = bmp.Encode(&NewBuf, set)

	case "tiff":
		err = tiff.Encode(&NewBuf, set, nil)

	default:
		return nil, errors.New("该图片格式不支持压缩")
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return NewBuf.Bytes(), nil
}

//获取icon
func GetIcon(pathSrc string) string {
	old, err := os.ReadFile(pathSrc)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var sizes = []uint{16, 24, 32, 48, 128, 256}
	logger.Log.Info(utils.PathExists("icon"))

	if !utils.PathExists("icon") {
		os.Mkdir("icon", 0666)
	}
	_, baseName := filepath.Split(pathSrc)
	ext := path.Ext(baseName)
	logger.Log.Info(baseName, ext)

	for _, size := range sizes {
		new_pic, err := CutImageResource(old, size, size, "png")
		if err != nil {
			logger.Log.Error(err)
			return ""
		}
		logger.Log.Info(fmt.Sprintf("icon/%s_%d.png", strings.TrimSuffix(baseName, ext), size))

		ioutil.WriteFile(fmt.Sprintf("icon/%s_%d.png", strings.TrimSuffix(baseName, ext), size), new_pic, 0666)
	}
	logger.Log.Info("获取icon成功!!!!!!!!!!")
	dirPath, _ := os.Getwd()

	return dirPath + "/icon/"
}

//修改大小
func Reset(filePath, targetPath string, width, height uint) string {
	old, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	new_pic, err := CutImageResource(old, width, height, "png")
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	_, baseName := filepath.Split(filePath)
	ext := path.Ext(filePath)
	name := strings.TrimSuffix(baseName, ext)
	if targetPath == "" {
		if !utils.PathExists("./reset") {
			os.Mkdir("./reset", 0666)
		}
		ioutil.WriteFile(fmt.Sprintf("./reset/%s.%s", name, ext), new_pic, 0666)
		fmt.Println("修改成功!!!!!!!!!!")
		dirPath, _ := os.Getwd()

		return dirPath + "/reset/"
	} else {
		if utils.IsDir(targetPath) {
			ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", targetPath, name, ext), new_pic, 0666)
			return targetPath
		} else {
			ioutil.WriteFile(targetPath, new_pic, 0666)
			paths, _ := filepath.Split(targetPath)
			return paths
		}
	}
}

//修改图片格式
func ChangePhoto(filePath, targetPath, layout string) string {
	old, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	im, _, err := image.Decode(bytes.NewReader(old))
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	_, baseName := filepath.Split(filePath)
	ext := path.Ext(filePath)
	name := strings.TrimSuffix(baseName, ext)
	width, height := im.Bounds().Dx(), im.Bounds().Dy()
	new_pic, err := CutImageResource(old, uint(height), uint(width), layout)
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	if targetPath == "" {
		if !utils.PathExists("./newPhoto") {
			os.Mkdir("./newPhoto", 0666)
		}
		logger.Log.Info(name, layout)

		ioutil.WriteFile(fmt.Sprintf("./newPhoto/%s.%s", name, layout), new_pic, 0666)
		dirPath, _ := os.Getwd()
		logger.Log.Info("修改成功!!!!!!!!!!")
		return dirPath + "/newPhoto/"
	} else {
		if utils.IsDir(targetPath) {
			ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", targetPath, name, layout), new_pic, 0666)
			return targetPath
		} else {
			ioutil.WriteFile(targetPath, new_pic, 0666)
			paths, _ := filepath.Split(targetPath)
			logger.Log.Info("修改成功!!!!!!!!!!")
			return paths
		}
	}

}
