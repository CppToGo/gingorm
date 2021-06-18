package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Benchmark_Router(b *testing.B){
	b.StopTimer()
	b.N = 1000
	b.StartTimer()
	for uid := 0; uid < b.N ; uid ++ {
		http.Get(fmt.Sprintf("http://localhost:8080/db/datainfo/%d", uid + 1000000))
	}
}