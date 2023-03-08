package server

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/lijianjunljj/gmfcommon/config"
	"github.com/lijianjunljj/gmfcommon/router"
	"golang.org/x/sync/errgroup"
)

type AbstractServer interface {
	BeforeRun(*config.Config) micro.Service
	Run(*config.Config) error
	GetName() string
	GetServiceName() string
	ServiceClient() interface{}
	RegisterServiceHandlerFunc() error
	GetWebRouter() router.AbstractRouter
	ErrGroup() *errgroup.Group
	EtcdReg(*config.EtcdOptions) registry.Registry
}
