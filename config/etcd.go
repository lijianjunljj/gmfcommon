package config

type EtcdOption func(o *EtcdOptions)
type EtcdOptions struct {
	RegistryAddr string
}

func NewEtcdOptions(opts ...EtcdOption) EtcdOptions {
	opt := EtcdOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
func RegistryAddr(v string) EtcdOption {
	return func(o *EtcdOptions) {
		o.RegistryAddr = v
	}
}
