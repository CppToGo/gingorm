package main

import "gingorm/src/Framwork"

type TestServer struct {
	Framwork.Service
}


func (s *TestServer) LoadConfig() error {
	s.Engine.LoadHTMLGlob("./template")
	return nil
}


func (s *TestServer) RegisterRouter() error {
	var err error
	err = s.RegisterPath(new(CoredataQueryHandle))
	return err
}
