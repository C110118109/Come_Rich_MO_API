package build_file

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func CreateExcel(output string) (path string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 設定儲存格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	saving := "storage" + string(os.PathSeparator) + output + ".xlsx"
	if err := f.SaveAs(saving); err != nil {
		fmt.Println(err)

	}
	path = "http://127.0.0.1:8090/public/" + output + ".xlsx"
	return path
}
