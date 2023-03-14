package config

type RedisOption func(o *RedisOptions)
type RedisOptions struct {
	Addr     string
	Timeout string
	Password string
	Maxidle string
	MaxActive string
	MaxIdleTimeout string
	DbcachePublic string
	DbauthAdmin string
	DbauthUser string
	DbcacheUser string
	DbcacheAdmin string
	DbFile string
}

func NewRedisOptions(opts ...RedisOption) RedisOptions {
	opt := RedisOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func RedisAddr(v string) RedisOption {
	return func(o *RedisOptions) {
		o.Addr = v
	}
}

func RedisTimeout(v string) RedisOption {
	return func(o *RedisOptions) {
		o.Timeout = v
	}
}

func RedisPassword(v string) RedisOption {
	return func(o *RedisOptions) {
		o.Password = v
	}
}

func RedisMaxidle(v string) RedisOption {
	return func(o *RedisOptions) {
		o.Maxidle = v
	}
}

func RedisMaxActive(v string) RedisOption {
	return func(o *RedisOptions) {
		o.MaxActive = v
	}
}

func RedisMaxIdleTimeout(v string) RedisOption {
	return func(o *RedisOptions) {
		o.MaxIdleTimeout = v
	}
}

func RedisDbcachePublic(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbcachePublic = v
	}
}

func RedisDbauthAdmin(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbauthAdmin = v
	}
}

func RedisDbauthUser(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbauthUser = v
	}
}

func RedisDbcacheUser(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbcacheUser = v
	}
}

func RedisDbcacheAdmin(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbcacheAdmin = v
	}
}

func RedisDbFile(v string) RedisOption {
	return func(o *RedisOptions) {
		o.DbFile = v
	}
}