package build_file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func OpenFile(outputName string) (path string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := filepath.Join(currentDir, "come_rich_model/come_rich_model.xlsx")

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 獲取工作表中指定儲存格的值
	cell, err := f.GetCellValue("CRM", "A7")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	// 獲取 Sheet1 上所有儲存格

	// rows, err := f.GetRows("PI")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	saving := "storage" + string(os.PathSeparator) + outputName + ".xlsx"
	if err := f.SaveAs(saving); err != nil {
		fmt.Println(err)

	}
	path = "http://127.0.0.1:8090/public/" + outputName + ".xlsx"
	return path
}
