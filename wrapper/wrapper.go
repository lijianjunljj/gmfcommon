package wrapper

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"time"
)

type HystrixConfigFunc func(*hystrix.CommandConfig)

var hystrixConfig = &hystrix.CommandConfig{
	Timeout:                int(3 * time.Second), /// 执行command的超时时间为3s
	RequestVolumeThreshold: 20,                   //10s内的请求数量，达到这个请求数量后才去判断是否要开启熔断
	ErrorPercentThreshold:  50,                   //错误百分比，请求数量大于等于RequestVolumeThreshold且错误率到达此百分比后就会启动熔断
	SleepWindow:            int(2 * time.Second), //当熔断器被打开后，多久后去监测服务是否可用
}

func NewHystrixConfig(opts ...HystrixConfigFunc) *hystrix.CommandConfig {
	for _, o := range opts {
		o(hystrixConfig)
	}
	return hystrixConfig
}
func Timeout(timeout int) HystrixConfigFunc {
	return func(c *hystrix.CommandConfig) {
		c.Timeout = timeout
	}
}

func RequestVolumeThreshold(requestVolumeThreshold int) HystrixConfigFunc {
	return func(c *hystrix.CommandConfig) {
		c.RequestVolumeThreshold = requestVolumeThreshold
	}
}

func ErrorPercentThreshold(errorPercentThreshold int) HystrixConfigFunc {
	return func(c *hystrix.CommandConfig) {
		c.ErrorPercentThreshold = errorPercentThreshold
	}
}

func SleepWindow(sleepWindow int) HystrixConfigFunc {
	return func(c *hystrix.CommandConfig) {
		c.SleepWindow = sleepWindow
	}
}

type Wrapper struct {
	client.Client
}

func (wrapper *Wrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	hystrix.ConfigureCommand(cmdName, *hystrixConfig)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewWrapper(c client.Client, opts ...HystrixConfigFunc) *Wrapper {
	NewHystrixConfig(opts...)
	return &Wrapper{c}
}
