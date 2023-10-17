package order

import (
	build_file "eirc.app/internal/pkg/build_file"
	"eirc.app/internal/pkg/code"
	//"eirc.app/internal/pkg/util"
)

func CreateExcel() interface{} {

	downloadPath := build_file.CreateExcel("comeRich_excel")

	//產生excel檔案 呼叫
	return code.GetCodeMessage(code.Successful, downloadPath)
}
