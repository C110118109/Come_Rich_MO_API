package gorm

import (
	//"fmt"
	//"os"

	//"strconv"

	//"eirc.app/internal/pkg/log"
	//"gorm.io/driver/postgres"

	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	//"gorm.io/plugin/dbresolver"
)

func New() (*gorm.DB, error) {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://ManuOrder:nkuste215@163.18.42.223:1433?database=XMLY5000"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func New() (*gorm.DB, error) {
// 	const config string = "host=%s port=%s user=%s dbname=%s sslmode=%s password=%s"
// 	sources := fmt.Sprintf(config,
// 		os.Getenv("SOURCES_HOST"),
// 		os.Getenv("SOURCES_PORT"),
// 		os.Getenv("SOURCES_USER"),
// 		os.Getenv("SOURCES_DATABASE"),
// 		os.Getenv("SOURCES_SSLMODE"),
// 		os.Getenv("SOURCES_PASSWORD"),
// 	)
// 	db, err := gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  sources,
// 		PreferSimpleProtocol: true,
// 		//取消條件約束綁定(自動建表)
// 	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
// 	if err != nil {
// 		return nil, err
// 	}
// 	dbReplicas, err := strconv.Atoi(os.Getenv("DB_REPLICAS"))
// 	if err != nil {
// 		log.Debug(err)
// 	}
// 	if dbReplicas >= 1 {
// 		var dialectics []gorm.Dialector
// 		for i := 1; i <= dbReplicas; i++ {
// 			replicas := fmt.Sprintf(config,
// 				os.Getenv("REPLICAS_HOST_"+strconv.Itoa(i)),
// 				os.Getenv("REPLICAS_PORT_"+strconv.Itoa(i)),
// 				os.Getenv("REPLICAS_USER_"+strconv.Itoa(i)),
// 				os.Getenv("REPLICAS_PASSWORD_"+strconv.Itoa(i)),
// 				os.Getenv("REPLICAS_DATABASE_"+strconv.Itoa(i)),
// 				os.Getenv("REPLICAS_SSLMODE_"+strconv.Itoa(i)),
// 			)
// 			director := postgres.New(postgres.Config{
// 				DSN:                  replicas,
// 				PreferSimpleProtocol: true,
// 			})
// 			dialectics = append(dialectics, director)
// 		}
// 		err = db.Use(dbresolver.Register(dbresolver.Config{Replicas: dialectics}))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 	}
// 	return db, nil
// }

func NewSQLite() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
