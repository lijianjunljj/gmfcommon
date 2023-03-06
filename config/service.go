package config

type ServiceOption func(o *ServiceOptions)
type ServiceOptions struct {
	ServiceName string
	Address     string
}

func Address(v string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Address = v
	}
}

func ServiceName(v string) ServiceOption {
	return func(o *ServiceOptions) {
		o.ServiceName = v
	}
}

func NewServiceOptions(opts ...ServiceOption) ServiceOptions {
	opt := ServiceOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
