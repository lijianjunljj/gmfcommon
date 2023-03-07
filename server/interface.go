package server

import (
	"github.com/lijianjunljj/gmfcommon/config"
	"github.com/lijianjunljj/gmfcommon/router"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"golang.org/x/sync/errgroup"
)

type AbstractServer interface {
	BeforeRun(*config.Config) micro.Service
	Run(*config.Config, ...interface{}) error
	GetName() string
	GetServiceName() string
	ServiceClient() interface{}
	RegisterServiceHandlerFunc() error
	GetWebRouter() router.AbstractRouter
	ErrGroup() *errgroup.Group
	EtcdReg(*config.EtcdOptions) registry.Registry
}
