package excel

import (
	"testing"
)

func TestExcel(t *testing.T) {
	excelObj := NewExcelObj()
	title := []string{"姓名", "年龄", "性别"}
	rows := [][]interface{}{
		{"张三", 18, "男"},
		{"李四", 19, "男"},
		{"Annie", 20, "女"},
	}
	excelObj.WriteExcel(title, rows)
	err := excelObj.file.SaveAs("./test.xlsx")
	if err != nil {
		t.Error(err)
		return
	}
}
