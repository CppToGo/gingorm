package main

import (
	"gingorm/Framwork"
)

func main() {

	fw := Framwork.Framework{}

	fw.Init(new(TestServer))

	fw.Run()
}
