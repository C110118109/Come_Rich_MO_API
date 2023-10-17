package build_file

import (
	"encoding/json"
	"fmt"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type ComeRich struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Body      struct {
		CustomerName    string    `json:"customer_name"`
		CustomerNo      string    `json:"customer_no"`
		SalesDate       time.Time `json:"sales_date"`
		SalesNo         string    `json:"sales_no"`
		InvoiceNumber   string    `json:"invoice_number"`
		DeliveryAddress string    `json:"delivery_address"`
		Total           int       `json:"total"`
		SalesRemark     string    `json:"sales_remark"`
		CustomerInfo    struct {
			No    string `json:"no"`
			Name  string `json:"name"`
			Tel   string `json:"tel"`
			Fax   string `json:"fax"`
			Unino string `json:"unino"`
		} `json:"CustomerInfo"`
		GoodsDetail []struct {
			ItemNum    int    `json:"item_num"`
			Name       string `json:"name"`
			Quantity   int    `json:"quantity"`
			Unit       string `json:"unit"`
			Price      int    `json:"price"`
			Subtotal   int    `json:"subtotal"`
			Remark     string `json:"remark"`
			Identifier string `json:"identifier"`
			FileID     string `json:"file_id"`
			ImgPath    string `json:"img_path"`
		} `json:"GoodsDetail"`
	} `json:"body"`
}

func DemoComeRich(outputName string) (path string) { //salesino input base

	//開啟前端資料
	url := "http://192.168.50.33:8090/come-rich/v1.0/sales-info/202305280003"

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwidHlwZSI6IkpXRSIsInppcCI6IkRFRiJ9.sdNxJKSFL9_-7mIPyZ1z1uaPuGHUYXgZm3ocrwZxqMyuKJEDruO9RRzOhhgw4401WVCXz1k8toEz-S8UQ0fEGduXS82tVl2RUfXJzXXYqfDseqJoq2oFQOAsK30giTtq1qP6PI3uZncUW7m8wo9-cujEkwIr1OpvLlYRMsSeLGE5r56vEIcvA16YHbpkdaaKbsC23s722ofJB4BmK6pFWsvNc2qcf8DetDcBwWNx6ci3bgKLdx3jZIZUw_4QdwypdUXV9sCG6PIMOvUFM5CFUTPdBTPF84dkIWyYaPgIehUE_N8I0BB9hNPTTUDZuOz-cXOEYpuC0YsQ6F9i3k1BKkCiLK_MFcAb0sHPmtHnw0EjPqZ4oycmX6Yt2MjCVwF3W-nYrFke8QNt9z7HUcP9_7Uv9qTiaCkcGVcc8h7r4aXEvFOiPItwRgi49noWD89O7xxLqOo2BxhK7QtmRflN4sp_xqoNVN7PtbJ4bhrJxU72Ne5tlyDpE6t2qmMJqat-lejkjMphMCfTOH6zQWkLu6Gbh_gtxoh-9W2FOePJsQj60pWEFAGZKkQURmXkt4MZnnz6iUI-oBfYTVy_5iR2h1KPyKoUN95WHGmopzWQednpoPjOxUW6CGeyjOt10Iz9RKup31st-jWARIPIrrRvrRzggbkolsBvgBjdimy0ttc.vMGTcskYYEgrikTr9gg7oQ.2xENjrh5KEv4aibQpDQJI46CIsnDuogQNFGbPf1b5valjrUjTkTGH-96p0bBVjlG8wCltu78ImUqLiVzdF2mDwiJ2BWzF2l6WguH8dOPPfBbWxzrUit2NX7Z9WrDvPnIzuLW5a5Z6AnPVS-TDuqlMtq6wft-Tomcr5wPMdqOdeCg4aXVZhc3M6EYxapfzyPJfpls6vIfYA7OAJdxB2Me0g.qM8q9EHiDwMDh3Tym3RwkB5DB8fvxOOSaetR7D0IH1A")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//反序列
	var readComeRich ComeRich
	json.Unmarshal(body, &readComeRich)
	//fmt.Println("A", readComeRich)
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
	GoodsDetailAmmount := len(readComeRich.Body.GoodsDetail) //資料數量
	GoodsDetailGap := GoodsDetailAmmount - 4                 //與實際表格差異
	if GoodsDetailAmmount > 4 {                              //模板是4格
		f.InsertRows("PI", 19, GoodsDetailGap) //插入新的欄位
		//f.DuplicateRow("PI",19)
	}

	for i := range readComeRich.Body.GoodsDetail {
		f.SetCellValue("PI", "A"+strconv.Itoa(16+i), readComeRich.Body.GoodsDetail[i].ItemNum) //寫下序號
		f.SetCellValue("PI", "B"+strconv.Itoa(16+i), readComeRich.Body.GoodsDetail[i])         //寫下品名
		f.SetCellValue("PI", "C"+strconv.Itoa(16+i), readComeRich.Body.GoodsDetail[i])         //寫下品名

	}

	//插入圖片
	pictureUrl := "http://192.168.50.208:4200/assets/images/%E9%A6%AC%E6%A8%99LOGO2.png"

	resp, _ := http.Get(pictureUrl)
	imgData, _ := io.ReadAll(resp.Body)
	enable, disable := true, false

	f.AddPictureFromBytes("PI", "E1", &excelize.Picture{
		Extension: ".png",
		File:      imgData,
		Format: &excelize.GraphicOptions{ScaleX: 0.3,
			ScaleY:          0.2,
			OffsetX:         30,
			OffsetY:         5,
			PrintObject:     &enable,
			LockAspectRatio: false,
			Locked:          &disable},
	})

	saving := "storage" + string(os.PathSeparator) + outputName + ".xlsx"
	if err := f.SaveAs(saving); err != nil {
		fmt.Println(err)

	}
	path = "http://127.0.0.1:8090/public/" + outputName + ".xlsx"
	return path
}
