package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"sort"
)

//excel 公共方法
const (
	WorkSheet      = "Sheet1"
	FirstRowHeight = 22
	RowHeight      = 20
	RowWidth       = 12
)

type ExcelInterface interface {
	ExcelDownload(titles []string, dataMap map[int][]interface{}) (dataByte []byte, err error)
}

type Excel struct{}

func NewExcelObj() *Excel {
	return &Excel{}
}

// excel下载
func (e *Excel) ExcelDownload(titles []string, dataMap map[int]*[]interface{}) (dataByte []byte, err error) {
	f := excelize.NewFile() //创建excel
	f.SetActiveSheet(0)     //设置默认工作表
	f.SetSheetRow(WorkSheet, "A1", &titles)
	//设置首行的高度
	f.SetRowHeight(WorkSheet, 1, FirstRowHeight)
	//写入数据
	lints := make([]int, 0)
	for k := range dataMap {
		lints = append(lints, k)
	}
	sort.Ints(lints)
	for index, lint := range lints {
		l := index + 2
		//行数据写入
		f.SetSheetRow(WorkSheet, fmt.Sprintf("A%d", l), dataMap[lint])
		f.SetRowHeight(WorkSheet, l, RowHeight)
	}
	//excel转二进制流
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
