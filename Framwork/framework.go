package Framwork

import (
	// "fmt"

	"fmt"
	"github.com/gin-gonic/gin"
)

type Framework struct {
	s ServiceInfc
}


func (fw *Framework)Init(s ServiceInfc){
	defer func() {
		if pic := recover() ; pic != nil {
			fmt.Println(pic)
		}
	}()
	var err error
	err  = s.Init()
	if err != nil {
		panic(err)
	}
	//加载配置
	err = s.LoadConfig()
	if err != nil {
		panic(err)
	}
	// 注册路由
	err = s.RegisterRouter()
	if err != nil {
		panic(err)
	}
	
	fw.s = s
	return
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
	h.RootRouter = "/"
}

func (h *Handle) Process(group *gin.RouterGroup) error {
	return nil
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
	// 获取实例
	s.Engine = gin.Default()
	return nil
}

func (s *Service) LoadConfig() error {
	return nil
}

func (s *Service) RegisterPath(handler HandleInfc) error {
	handler.SetRootRouter()
	g1 := s.Engine.Group(handler.GetRootRouter())
	return handler.Process(g1)
}

func (s *Service) RegisterRouter() error {
	var err error
	err = s.RegisterPath(new(Handle))
	return err
}

func (s *Service) Run() error {
	return s.Engine.Run(s.ListenAddr)
}
