package build_file

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	sales_info "eirc.app/internal/v1/structure/sales_info"
	"github.com/xuri/excelize/v2"

	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

func BuildComrichS3(outputName string, readContent sales_info.Base) (fileName string, filePath string, version string) {

	ctx := context.TODO()
	client := NewS3Client(ctx)

	bucket := "beibei14" //空間 bucket名稱
	key := "template/comeRich_excel.xlsx"
	in := CreateGetObjectInput(bucket, key)
	content := GetObjectContent(ctx, client, in)
	fmt.Println("S3載入 PI模板資料讀取成功") //模板資料請透過主控台進去建

	//f, err := excelize.OpenFile("./storage/piModle.xlsx")
	f, err := excelize.OpenReader(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 計算實際商品數量是否超過模板表格
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

	//上傳檔案 to S3
	buf, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}

	//存入S3路徑指定
	fileName = outputName + ".xlsx"
	input := CreateInput(bucket, "excel_file/PI/"+fileName, strings.NewReader((buf.String())))

	uploader := manager.NewUploader(client)
	status, err := uploader.Upload(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
	}
	json.Marshal(status)
	//fmt.Println(string(marshal))
	filePath = status.Location
	version = string(*status.VersionID)

	fmt.Println(filePath)
	fmt.Println(version)

	//回傳連結為S3連結(請務必設定2008公開政策)&版本
	return fileName, filePath, version
}
