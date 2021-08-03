package xmindgo

import (
	"fmt"

	"github.com/xiaobing94/xmindgo/models"
)

func (f *File) ActiveSheet(sheetId string) error {
	sheets := f.Workbook.GetSheets()
	var isFoundSheet bool
	for _, sheet := range sheets {
		if sheet.ID == sheetId {
			isFoundSheet = true
			break
		}
	}
	if !isFoundSheet {
		return fmt.Errorf("can not found sheetId:%s", sheetId)
	}
	if f.MetaData == nil {
		f.MetaData = models.NewMetaData()
	}
	f.MetaData.ActiveSheetId = sheetId
	return nil
}

func (f *File) SetCreator(name string) {
	if f.MetaData == nil {
		f.MetaData = models.NewMetaData()
	}
	f.MetaData.Creator.Name = name
}

func (f *File) SetVersion(version string) {
	if f.MetaData == nil {
		f.MetaData = models.NewMetaData()
	}
	f.MetaData.Creator.Version = version
}
