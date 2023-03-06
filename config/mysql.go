package config

type MysqlOption func(o *MysqlOptions)
type MysqlOptions struct {
	Db               string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassWord       string
	DbName           string
	MysqlTimeout     string
	MysqlLifeTimeout int
	MysqlMaxOpenCons int
	MysqlMaxIdleCons int
}

func NewMysqlOptions(opts ...MysqlOption) MysqlOptions {
	opt := MysqlOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func MysqlMaxIdleCons(v int) MysqlOption {
	return func(o *MysqlOptions) {
		o.MysqlMaxIdleCons = v
	}
}

func MysqlMaxOpenCons(v int) MysqlOption {
	return func(o *MysqlOptions) {
		o.MysqlMaxOpenCons = v
	}
}

func MysqlLifeTimeout(v int) MysqlOption {
	return func(o *MysqlOptions) {
		o.MysqlLifeTimeout = v
	}
}

func MysqlTimeout(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.MysqlTimeout = v
	}
}

func Db(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.Db = v
	}
}

func DbHost(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DbHost = v
	}
}

func DbPort(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DbPort = v
	}
}

func DbUser(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DbUser = v
	}
}

func DbPassWord(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DbPassWord = v
	}
}

func DbName(v string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DbName = v
	}
}

//func (mysql *MysqlOptions) Parse(filePath string) {
//
//}
//
//func (mysql *MysqlOptions) Init() {
//
//}
