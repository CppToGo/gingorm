package main

import (
	"gingorm/src/Framwork"
)

func main() {

	fw := Framwork.Framework{}
	server := new(TestServer)
	_ = fw.Init(server)

	_ = fw.Run()
}
