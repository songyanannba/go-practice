package helper

import (
	"github.com/xuri/excelize/v2"
)

func FastExportCol[V comparable](fileName string, cell string, arr []V) {
	f := excelize.NewFile()
	defer f.Close()
	f.SetSheetCol("Sheet1", cell, &arr)
	f.SaveAs(fileName)
}

func FastReadCol(fileName string, cellNum int) (data []string, err error) {
	f, err := excelize.OpenFile(fileName)
	defer f.Close()
	if err != nil {
		return
	}
	cols, err := f.GetCols("Sheet1")
	if err != nil {
		return
	}
	data = SliceVal(cols, cellNum-1)
	return
}

//
//type excel struct {
//	File *excelize.File
//}
//
//func NewExcel() *excel {
//	return &excel{
//		File: excelize.NewFile(),
//	}
//}
//
//func (e *excel) SetFields(fields []string) {
//	e.File.SetSheetRow("Sheet1", "A1", &fields)
//}
//
//func (e *excel) SetData(data []map[string]interface{}) {
//	var i = 0
//	for _, d := range data {
//		var arr []interface{}
//		for _, v := range d {
//			arr = append(arr, v)
//		}
//		e.File.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+2), &arr)
//		i++
//	}
//}
//
//func (e *excel) Save(path string) error {
//	return e.File.SaveAs(path)
//}
//
//func Export(data []map[string]interface{}) {
//	e := NewExcel()
//	e.SetFields([]string{"id", "name"})
//	e.SetData(data)
//	e.Save("test.xlsx")
//}
