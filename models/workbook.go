package models

import (
	"errors"
	"fmt"
)

type Workbook struct {
	sheets []*Sheet
}

// AddSheet 在Workbook添加一个sheet
func (w *Workbook) AddSheet(sheet *Sheet) {
	w.sheets = append(w.sheets, sheet)
}

// SetSheets 在Workbook设置sheets
func (w *Workbook) SetSheets(sheets []*Sheet) {
	w.sheets = sheets
}

// RemoveSheet 移除workbook中的sheet
func (w *Workbook) RemoveSheet(sheetID string) {
	sheets := w.sheets
	k := 0
	for i, sheet := range w.sheets {
		if sheet.ID == sheetID {
			if i != k {
				sheets[k] = sheet
			}
			k ++
		}
	}
	w.sheets = sheets[:k]
}

// CreateEmptySheet 创建空的Sheet
func (w *Workbook) CreateEmptySheet() *Sheet {
	sheet := &Sheet{}
	sheet.GenID()
	w.sheets = append(w.sheets, sheet)
	return sheet
}

func (w *Workbook) GetSheetByIndex(index int)(*Sheet, error) {
	if len(w.sheets) <= index {
		return nil, errors.New("out of range index")
	}
	return w.sheets[index], nil
}

func (w *Workbook) GetSheets() []*Sheet {
	return w.sheets
}

func (w *Workbook) GetSheetByID(sheetID string) (*Sheet, error) {
	for _, sheet := range w.sheets {
		if sheet.ID == sheetID {
			return sheet, nil
		}
	}
	return nil, fmt.Errorf("sheet index %s not found", sheetID)
}