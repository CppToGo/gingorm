package main

import (
	"gingorm/src/Framwork"
)

func main() {

	fw := Framwork.Framework{}

	fw.Init(new(TestServer))

	fw.Run()
}
