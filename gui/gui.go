package gui

// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"os"
	"photo/change"
	"photo/logger"
	"photo/utils"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type MyMainWindow struct {
	*walk.MainWindow
	SrcPathedit        *walk.LineEdit
	TargetPathedit     *walk.LineEdit
	TargetWidthedit    *walk.LineEdit
	TargetHeightedit   *walk.LineEdit
	ChooseComboBox     *walk.ComboBox
	TargetTypeComboBox *walk.ComboBox
	webView            *walk.WebView
}
type myStruct struct {
	Choose       int
	SrcPath      string
	TargetPath   string
	TargetWidth  uint
	TargetHeight uint
	TargetType   string
}
type Species struct {
	Id   int
	Name string
}
type Options struct {
	Label string
	Value string
}

var my = &myStruct{
	Choose:       1,
	SrcPath:      "",
	TargetWidth:  0,
	TargetHeight: 0,
	TargetType:   "png",
}

const (
	SIZE_W = 900
	SIZE_H = 500
)

func Gui() {
	logger.InitLog(true)

	mw := new(MyMainWindow)
	var db *walk.DataBinder

	err := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "图片处理小工具",
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "myStruct",
			DataSource:     my,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Icon:   "icon.png",
		Size:   Size{900, 500},
		Layout: VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					VSplitter{
						MaxSize: Size{350, 200},
						Children: []Widget{HSplitter{
							Children: []Widget{
								Label{Text: "选择方式: "},
								ComboBox{
									AssignTo:      &mw.ChooseComboBox,
									Value:         Bind("Choose"),
									BindingMember: "Id",
									DisplayMember: "Name",
									Model: []*Species{
										{1, "生成小图标"},
										{2, "修改图片大小"},
										{3, "图片文件转换"},
									},
								},
							},
						},
							HSplitter{
								Children: []Widget{
									Label{
										Text: "源文件:",
									},
									LineEdit{
										Text:     Bind("SrcPath"),
										AssignTo: &mw.SrcPathedit,
									},
									//普通按钮
									PushButton{
										Text:      "选择文件或目录",
										OnClicked: mw.selectFile, //点击事件响应函数
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Text: "目标路径:",
									},
									LineEdit{
										Text:     Bind("TargetPath"),
										AssignTo: &mw.TargetPathedit,
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Text: "目标图片宽度:",
									},
									LineEdit{
										Text:     Bind("TargetWidth"),
										AssignTo: &mw.TargetWidthedit,
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Text: "目标图片高度:",
									},
									LineEdit{
										Text:     "",
										AssignTo: &mw.TargetHeightedit,
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Text: "转换类型:",
									},
									ComboBox{
										AssignTo:      &mw.TargetTypeComboBox,
										Value:         Bind("TargetType"),
										BindingMember: "Value",
										DisplayMember: "Label",
										Model: []*Options{
											{"png", "png"},
											{"jpg", "jpg"},
											{"jpeg", "jpeg"},
											{"webp", "webp"},
											{"webp2", "webp2"},
											{"bmp", "bmp"},
											{"tiff", "tiff"},
										},
									},
								},
							},
							//普通按钮
							PushButton{
								Text:      "完成",
								OnClicked: mw.sumbit, //点击事件响应函数
							}},
					},
					WebView{
						AssignTo:      &mw.webView,
						StretchFactor: 2,
					},
				},
			},
		},
	}.Create() //创建
	defaultStyle := win.GetWindowLong(mw.Handle(), win.GWL_STYLE) // Gets current style
	newStyle := defaultStyle &^ win.WS_THICKFRAME                 // Remove WS_THICKFRAME
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, newStyle)

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	if err != nil {
		logger.Log.Error(err)
		os.Exit(1)
	}

	mw.Run() //运行
}

func (mw *MyMainWindow) selectFile() {

	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "选择文件 (*.webp;*.webp2;*.bmp;*.jpeg;*.jpg;*.png;*.tiff)|*.webp;*.webp2;*.bmp;*.jpeg;*.jpg;*.png;*.tiff"

	mw.SrcPathedit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.SrcPathedit.SetText("Error : File Open\r\n")
		return
	} else if !ok {
		return
	}
	filePath := dlg.FilePath
	mw.SrcPathedit.SetText(filePath)
}
func (mw *MyMainWindow) sumbit() {
	if mw.SrcPathedit.Text() == "" {
		logger.Log.Error("源地址不能为空")
		Popup2(mw, "源地址不能为空!!!!!!!!!!")
		return
	}
	my = &myStruct{
		Choose:       mw.ChooseComboBox.CurrentIndex(),
		TargetType:   mw.TargetTypeComboBox.Text(),
		SrcPath:      mw.SrcPathedit.Text(),
		TargetPath:   mw.TargetPathedit.Text(),
		TargetWidth:  utils.String2Number(mw.TargetWidthedit.Text()),
		TargetHeight: utils.String2Number(mw.TargetHeightedit.Text()),
	}
	if my.Choose == 0 {
		var p string
		if utils.IsDir(my.SrcPath) {
			//遍历出所有图片
			fileList := utils.GetImgFiles(my.SrcPath, "")
			for _, SrcPath := range fileList {
				p = change.GetIcon(SrcPath)
			}
		} else {
			p = change.GetIcon(my.SrcPath)
		}
		Popup2(mw, "生成成功!!!!!!!!!!")
		mw.webView.SetURL(p)
	} else if my.Choose == 1 {
		p := change.Reset(my.SrcPath, my.TargetPath, my.TargetWidth, my.TargetHeight)
		Popup2(mw, "修改成功!!!!!!!!!!")
		mw.webView.SetURL(p)
	} else if my.Choose == 2 {
		p := change.ChangePhoto(my.SrcPath, my.TargetPath, my.TargetType)
		Popup2(mw, "修改成功!!!!!!!!!!")
		mw.webView.SetURL(p)
	} else {
		logger.Log.Info("识别错误")
		Popup2(mw, "识别错误!!!!!!!!!!")
	}
}
func Popup2(mw *MyMainWindow, str string) {
	// 按键结果以int形式返回
	// cmd :=walk.MsgBox(
	walk.MsgBox(
		mw,
		"提示",
		str,
		walk.MsgBoxOK,
	)
}
