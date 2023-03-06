package config

type WebOption func(o *WebOptions)
type WebOptions struct {
	Addr     string
	Protocol string
}

func NewWebOptions(opts ...WebOption) WebOptions {
	opt := WebOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
func Protocol(v string) WebOption {
	return func(o *WebOptions) {
		o.Protocol = v
	}
}

func Addr(v string) WebOption {
	return func(o *WebOptions) {
		o.Addr = v
	}
}
