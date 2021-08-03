package main

import (
	"fmt"
	"io/ioutil"

	"github.com/xiaobing94/xmindgo"
	"github.com/xiaobing94/xmindgo/models"
)

func main() {
	// 生成xmind示例
	xmindFile := xmindgo.NewFile()
	sheet := xmindFile.Workbook.CreateEmptySheet()
	sheet.UseDefaultTheme()
	xmindFile.Workbook.AddSheet(sheet)
	topic := models.NewTopic("rootTopicTitle")
	sheet.AddRootTopic(topic)
	imageTopic := models.NewTopic("xmind")

	filename := "xmind.jpg"
	xmindFilename := "example.xmind"
	localImage, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	imageTopic.Image, err = xmindFile.CreateImage(localImage, filename)
	if err != nil {
		panic(err)
	}
	topic.AddAttachedChildTopic(imageTopic)
	if err := xmindFile.SaveAs("example.xmind"); err != nil {
		panic(err)
	}

	// 解析示例
	file, err := xmindgo.OpenFile(xmindFilename)
	if err != nil {
		panic(err)
	}
	workbook := file.GetWorkbook()
	sheet, err = workbook.GetSheetByIndex(0)
	if err != nil {
		panic(err)
	}
	// rootTopicTitle
	fmt.Println(sheet.RootTopic.Title)
}
