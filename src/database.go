package main

import (
	// "fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "time"
)

var DataBase *gorm.DB

func init() {
	// dsn := "root:123456@tcp(linux102:3306)/coredata?charset=utf8mb4&loc=Local"
	// var err error
	// DataBase ,  err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(fmt.Errorf("Link database failed err:%v ", err))
	// }
	// db, err := DataBase.DB()
	// if err != nil {
	// 	panic(fmt.Errorf("init db error:%v", err))
	// }
	// db.SetMaxIdleConns(512)
	// db.SetMaxOpenConns(512)
	// db.SetConnMaxLifetime(5 * time.Second)
}