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
	s.RegisterRouter()
	
	fw.s = s
	return
}

func (fw *Framework)Run()error{
	return fw.s.Run()
}

type GroupRouterInfc interface {
	Process(group *gin.RouterGroup)error
}

type GroupRouter struct {
}

func (h *GroupRouter) Process(group *gin.RouterGroup) error {
	panic("implement me")
}



type ServiceInfc interface {
	Init()error
	LoadConfig()error
	RegisterPath(PrefixRouter string, handler GroupRouterInfc)
	RegisterRouter()
	Run()error
}

type Service struct {
	Engine *gin.Engine
	Addr   string
}

func (s *Service) LoadConfig() error {
	panic("implement me")
}

func (s *Service) RegisterRouter(){
	panic("implement me")
}

func (s *Service) Init() error {
	// 获取实例
	s.Engine = gin.Default()
	// 加载deploy配置
	
	return nil
}

func (s *Service) RegisterPath(PrefixRouter string,handler GroupRouterInfc)  {
	group := s.Engine.Group(PrefixRouter)
	if group == nil {
		panic(fmt.Errorf("create group router failed: %v", PrefixRouter))
	}
	err := handler.Process(group)
	if err != nil {
		panic(fmt.Errorf("RegistProcess err :%v", err))
	}
}

func (s *Service) Run() error {
	return s.Engine.Run(s.Addr)
}
