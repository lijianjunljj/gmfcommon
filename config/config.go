package config

import (
	"gmf/src/common/config/parser"
)

type Config struct {
	Mysql   *MysqlOptions
	Etcd    *EtcdOptions
	Web     *WebOptions
	Service *ServiceOptions
	parser  parser.Parser
	DbType  string
}

func NewConfig(parser parser.Parser) *Config {
	//fmt.Println("parser", parser)
	config := &Config{
		parser: parser,
	}
	return config
}
func (c *Config) InitEtcd() {
	registryAddr := c.parser.GetString("etcd", "registryAddr")
	etcdConf := NewEtcdOptions(RegistryAddr(registryAddr))
	c.Etcd = &etcdConf
}
func (c *Config) InitService(serverName string) {
	serviceName := c.parser.GetString("servers", serverName, "serviceName")
	address := c.parser.GetString("servers", serverName, "address")
	serviceConf := NewServiceOptions(ServiceName(serviceName), Address(address))
	c.Service = &serviceConf
}

func (c *Config) InitWeb() {
	protocol := c.parser.GetString("gateway", "protocol")
	addr := c.parser.GetString("gateway", "addr")
	webConf := NewWebOptions(Protocol(protocol), Addr(addr))
	c.Web = &webConf
}
func (c *Config) InitDbType() {
	c.DbType = c.parser.GetString("db_type")
}
func (c *Config) InitMysql() {
	//fmt.Println("c.parser", c.parser)
	db := c.parser.GetString("mysql", "Db")
	//fmt.Println("db", db)
	dbHost := c.parser.GetString("mysql", "DbHost")
	dbPort := c.parser.GetString("mysql", "DbPort")
	dbName := c.parser.GetString("mysql", "DbName")
	dbUser := c.parser.GetString("mysql", "DbUser")
	dbPassWord := c.parser.GetString("mysql", "DbPassWord")
	mysqlTimeout := c.parser.GetString("mysql", "MysqlTimeout")
	mysqlLifeTimeout := c.parser.GetInt("mysql", "MysqlLifeTimeout")
	mysqlMaxOpenCons := c.parser.GetInt("mysql", "MysqlMaxOpenCons")
	mysqlMaxIdleCons := c.parser.GetInt("mysql", "MysqlMaxIdleCons")
	mysqlConf := NewMysqlOptions(Db(db), DbHost(dbHost), DbPort(dbPort),
		DbUser(dbUser), DbPassWord(dbPassWord), DbName(dbName), MysqlTimeout(mysqlTimeout),
		MysqlLifeTimeout(mysqlLifeTimeout), MysqlMaxOpenCons(mysqlMaxOpenCons), MysqlMaxIdleCons(mysqlMaxIdleCons),
	)
	//fmt.Println("mysqlConf:", mysqlConf)
	c.Mysql = &mysqlConf
}
