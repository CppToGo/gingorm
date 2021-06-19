package main

import (
	"gingorm/src/Framwork"
)

func main() {
	fw := Framwork.Framework{}
	server := new(TestServer)
	err := fw.Init(server)
	if err != nil {
		panic(err)
	}

	err = fw.Run()
	if err != nil{
		panic(err)
	}
}
