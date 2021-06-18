package Framwork

import (
	"github.com/gin-gonic/gin"
)

type Framework struct {
	s ServiceInfc
}


func (fw *Framework)Init(s ServiceInfc)error{
	fw.s = s
	return fw.s.Init()
}

func (fw *Framework)Run()error{
	return fw.s.Run()
}

type HandleInfc interface {
	GetRootRouter()string
	SetRootRouter()
	Process(group *gin.RouterGroup)error
}

type Handle struct {
	RootRouter string
}

func (h *Handle) SetRootRouter() {
	panic("implement me")
}

func (h *Handle) Process(group *gin.RouterGroup) error {
	panic("implement me")
}

func (h *Handle) GetRootRouter() string {
	return h.RootRouter
}



type ServiceInfc interface {
	Init()error
	LoadConfig()error
	RegisterPath(handler HandleInfc)error
	RegisterRouter()error
	Run()error
}

type Service struct {
	Engine     *gin.Engine
	ListenAddr string
}

func (s *Service) Init() error {
	var err error
	// 获取实例
	s.Engine = gin.Default()

	//加载配置
	err = s.LoadConfig()
	if err != nil {
		return err
	}
	// 注册路由
	err = s.RegisterRouter()
	if err != nil {
		return err
	}

	//启动
	err = s.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoadConfig() error {
	return nil
}

func (s *Service) RegisterPath(handler HandleInfc) error {
	g1 := s.Engine.Group(handler.GetRootRouter())
	return handler.Process(g1)
}

func (s *Service) RegisterRouter() error {
	panic("implement me")
}

func (s *Service) Run() error {
	return s.Engine.Run(s.ListenAddr)
}
