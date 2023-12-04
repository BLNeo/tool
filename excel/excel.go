package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// excel 公共方法
const (
	WorkSheet      = "Sheet1"
	FirstRowHeight = 22
	RowHeight      = 20
	RowWidth       = 12
)

type ExcelInterface interface {
	ExcelDownload(titles []string, dataMap map[int][]interface{}) (dataByte []byte, err error)
}

type Excel struct {
	file *excelize.File
}

func NewExcelObj() *Excel {
	return &Excel{
		excelize.NewFile(),
	}
}

// WriteExcel 写入表格中
func (e *Excel) WriteExcel(titles []string, rows [][]interface{}) {
	e.file.SetActiveSheet(0) //设置默认工作表
	e.file.SetSheetRow(WorkSheet, "A1", &titles)
	//设置首行的高度
	e.file.SetRowHeight(WorkSheet, 1, FirstRowHeight)
	//写入数据
	for index, row := range rows {
		l := index + 2
		//行数据写入
		e.file.SetSheetRow(WorkSheet, fmt.Sprintf("A%d", l), &row)
		e.file.SetRowHeight(WorkSheet, l, RowHeight)
	}
}

// ToBytes excel转二进制流
func (e *Excel) ToBytes() (dataByte []byte, err error) {
	buf, err := e.file.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
