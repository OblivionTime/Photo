/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-06-14 17:28:30
 * @LastEditors: solid
 * @LastEditTime: 2022-06-15 17:50:37
 */
package main

import (
	"flag"
	"fmt"
	"photo/change"
	"photo/gui"
)

var (
	srcPath           string
	targetPath        string
	generateSmallIcon bool
	modifySize        bool
	convertImage      bool
	targetWidth       uint
	targetHeight      uint
	targetType        string
)

func main() {
	gui.Gui()
	flag.Parse()
	if srcPath == "" {
		fmt.Println("原图片的路径不能为空")
		return
	}
	if generateSmallIcon {
		change.GetIcon(srcPath)
	} else if modifySize {
		change.Reset(srcPath, targetPath, targetWidth, targetHeight)
	} else if convertImage {
		change.ChangePhoto(srcPath, targetPath, targetType)
	} else {
		fmt.Println("识别错误")
	}

}
func init() {
	flag.StringVar(&srcPath, "src", "", "要修改图片的路径(必填)")
	flag.StringVar(&targetPath, "target", "", "生成的图片存放的路径(如果不填会自动生成到相对应位置)")
	flag.BoolVar(&generateSmallIcon, "g", false, "将png生成16,24,32,48,128,256大小的小图标")
	flag.BoolVar(&modifySize, "m", false, "修改原图片的尺寸")
	flag.BoolVar(&convertImage, "c", false, "原图片转换")
	flag.UintVar(&targetWidth, "w", 0, "目标图片的宽度")
	flag.UintVar(&targetHeight, "h", 0, "目标图片的高度")
	flag.StringVar(&targetType, "t", "jpg", "转换成目标的类型(png/jpg/jpeg/webp/webp2/bmp/tiff)")
}
