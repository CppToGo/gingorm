package main

import (
	"gingorm/src/Framwork"
)

func main() {
	fw := Framwork.Framework{}

	err := fw.Init(new(TestServer))
	if err != nil {
		panic(err)
	}

	err = fw.Run()
	if err != nil{
		panic(err)
	}
}
