package build_file

import (
	//"encoding/json"
	"fmt"
	//_ "image/png"
	"os"
	"path/filepath"
	"strconv"

	//"time"

	sales_info "eirc.app/internal/v1/structure/sales_info"
	"github.com/xuri/excelize/v2"
)

func DemoComeRich(outputName string, readContent sales_info.Base) (path string) {

	//開啟相對位置
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

	//計算實際商品數量是否超過模板表格
	GoodsDetailAmmount := len(readContent.GoodsDetail) //資料數量
	GoodsDetailGap := GoodsDetailAmmount - 4           //與實際表格差異
	if GoodsDetailAmmount > 4 {                        //模板是4格
		f.InsertRows("CRM", 19, GoodsDetailGap) //插入新的欄位
		//f.DuplicateRow("CRM",19)
	}

	for i := range readContent.GoodsDetail {
		f.SetCellValue("CRM", "A"+strconv.Itoa(16+i), readContent.GoodsDetail[i].ItemNum)  //寫下序號
		f.SetCellValue("CRM", "B"+strconv.Itoa(16+i), readContent.GoodsDetail[i].Name)     //寫下品名
		f.SetCellValue("CRM", "E"+strconv.Itoa(16+i), readContent.GoodsDetail[i].Quantity) //寫下品名
		f.SetCellValue("CRM", "G"+strconv.Itoa(16+i), readContent.GoodsDetail[i].Unit)     //寫下品名
		//f.SetCellValue("CRM", "E"+strconv.Itoa(16+i), readContent.GoodsDetail[i].Quantity) //寫下品名

	}

	// //插入圖片
	// pictureUrl := "http://192.168.50.208:4200/assets/images/%E9%A6%AC%E6%A8%99LOGO2.png"

	// resp, _ := http.Get(pictureUrl)
	// imgData, _ := io.ReadAll(resp.Body)
	// enable, disable := true, false

	// f.AddPictureFromBytes("CRM", "E1", &excelize.Picture{
	// 	Extension: ".png",
	// 	File:      imgData,
	// 	Format: &excelize.GraphicOptions{ScaleX: 0.3,
	// 		ScaleY:          0.2,
	// 		OffsetX:         30,
	// 		OffsetY:         5,
	// 		PrintObject:     &enable,
	// 		LockAspectRatio: false,
	// 		Locked:          &disable},
	// })

	saving := "storage" + string(os.PathSeparator) + outputName + ".xlsx"
	if err := f.SaveAs(saving); err != nil {
		fmt.Println(err)

	}
	// if err := f.SaveAs(outputName + ".xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
	path = "http://127.0.0.1:8090/public/" + outputName + ".xlsx"
	return path
}
