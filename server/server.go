package server

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/lijianjunljj/gmfcommon/config"
	commondb "github.com/lijianjunljj/gmfcommon/db"
	"github.com/lijianjunljj/gmfcommon/router"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	AbstractServer
	G               *errgroup.Group
	ServiceCallFunc func(microService micro.Service)
	Config          *config.Config
	Name            string
	ServiceName     string
	ClientService   micro.Service
	Service         interface{}
	WebRouter       router.AbstractRouter
	AutoAutoMigrateTables   []interface{}
}

func (s *Server) BeforeRun(config *config.Config) micro.Service {
	s.Config = config
	s.Config.InitService(s.Name)
	fmt.Println("config1111:", s.Config)
	// etcd注册件
	etcdReg := s.EtcdReg(s.Config.Etcd)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(s.ServiceName), // 微服务名字
		micro.Address(s.Config.Service.Address),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()

	return microService
}
func (s *Server) GetWebRouter() router.AbstractRouter {
	return s.WebRouter
}

func (s *Server) Run(config *config.Config) error {
	commondb.Init(config.DbType, func() interface{} {
		return config.Mysql
	}, s.AutoAutoMigrateTables...)
	microService := s.BeforeRun(config)
	// 服务注册
	if s.ServiceCallFunc != nil {
		s.ServiceCallFunc(microService)
	}

	// 启动微服务
	_ = microService.Run()
	return nil
}
func (s *Server) ErrGroup() *errgroup.Group {
	return s.G
}
func (s *Server) GetConfig() *config.Config {
	return s.Config
}
func (s *Server) GetName() string {
	return s.Name
}

func (s *Server) GetServiceName() string {
	return s.ServiceName
}

func (s *Server) ServiceClient() interface{} {
	return s.Service

}

func (s *Server) EtcdReg(options *config.EtcdOptions) registry.Registry {

	etcdReg := etcd.NewRegistry(
		registry.Addrs(options.RegistryAddr),
	)
	return etcdReg
}

var sm Manager

func Init() *Manager {
	sm.Init()
	return &sm
}

func StartAll(config *config.Config) {
	sm.RunAll(config)
}
func Start(name string, config *config.Config) {
	err := sm.Run(name, config)
	if err != nil {
		fmt.Println(err.Error())
	}

}
