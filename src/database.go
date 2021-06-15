package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func init() {
	dsn := "root:123456@tcp(linux102:3306)/coredata?charset=utf8mb4&loc=Local"
	var err error
	DataBase ,  err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Link database failed err:%v ", err))
	}
}