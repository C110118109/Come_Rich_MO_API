package main

// Project Author: Shane, shane871112@hotmail.com
// GCC require!! https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe
import (
	"net/http"

	"eirc.app/internal/pkg/dao/gorm"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/v1/router"
	routerAccount "eirc.app/internal/v1/router/account"
	routerCustomer "eirc.app/internal/v1/router/customer"
	routerFile "eirc.app/internal/v1/router/file"
	routerLogin "eirc.app/internal/v1/router/login"
	routerSalesInfo "eirc.app/internal/v1/router/sales_info"

	accountModel "eirc.app/internal/v1/structure/accounts"
	fileModel "eirc.app/internal/v1/structure/file"
)

// @version 0.1
// @author Shane
// @description COME RICH 製令平台

func main() {
	dbLY, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	db, err := gorm.NewSQLite()
	if err != nil {

		log.Error(err)
		return
	}
	db.AutoMigrate(&fileModel.Table{})
	db.AutoMigrate(&accountModel.Table{})

	route := router.Default()
	route = routerCustomer.GetRoute(route, dbLY)      //客戶路由
	route = routerSalesInfo.GetRoute(route, dbLY, db) //銷貨單路由
	route = routerFile.GetRoute(route, db, dbLY)      //檔案上傳路由
	route = routerAccount.GetRoute(route, db)
	route = routerLogin.GetRoute(route, db)

	log.Fatal(http.ListenAndServe("163.18.18.42:8090", route)) //localhost
}
