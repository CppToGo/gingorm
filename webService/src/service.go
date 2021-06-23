package main

import (
	"gingorm/Framwork"
)

type TestServer struct {
	Framwork.Service
}

func (s *TestServer) LoadConfig() error {
	s.Engine.LoadHTMLGlob("./template/*")
	s.ListenAddr = "localhost:8080"
	return nil
}


func (s *TestServer) RegisterRouter() error {
	var err error
	s.RegisterPath(new(CoredataQueryHandle))
	s.RegisterPath(new(TestHandle))
	return err
}