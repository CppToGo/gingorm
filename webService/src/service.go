package main

import (
	"gingorm/Framwork"
)

type TestServer struct {
	Framwork.Service
}

func (s *TestServer) LoadConfig() error {
	s.Engine.LoadHTMLGlob("./template/*")
	s.Addr = "localhost:8080"
	return nil
}


func (s *TestServer) RegisterRouter(){
	s.RegisterPath("/db", new(CoredataQueryHandle))
	s.RegisterPath("/test", new(TestHandle))
}