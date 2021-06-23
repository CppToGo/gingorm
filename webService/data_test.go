package main

import (
	"fmt"
	"log"
	"testing"
)

func Benchmark_Router(b *testing.B){
	b.StopTimer()
	b.StartTimer()
	for UID := 10030125 ; UID - 10030125 < 1 ; UID++{
		Index1 := UID % 512 + 1
		tableName := fmt.Sprintf("tbl_str_index_%03d",Index1)
		result := make([]StrResult, 10)
		dbrl := DataBase.Table(tableName).Where("dataid = ? and index1 = ?", 10003, UID).Find(&result)
		if dbrl.Error != nil {
			log.Fatal(dbrl)
		}
	}
}