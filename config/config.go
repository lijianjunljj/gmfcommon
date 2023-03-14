package config

import (
	"github.com/lijianjunljj/gmfcommon/config/parser"
)

type Config struct {
	Mysql   *MysqlOptions
	Etcd    *EtcdOptions
	Web     *WebOptions
	Redis   *RedisOptions
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
func (c *Config) GetString(opts ...string) string {
	return c.parser.GetString(opts...)
}

func (c *Config) InitEtcd() *Config {
	registryAddr := c.parser.GetString("etcd", "registryAddr")
	etcdConf := NewEtcdOptions(RegistryAddr(registryAddr))
	c.Etcd = &etcdConf
	return c
}
func (c *Config) InitService(serverName string) *Config {
	serviceName := c.parser.GetString("servers", serverName, "serviceName")
	address := c.parser.GetString("servers", serverName, "address")
	serviceConf := NewServiceOptions(ServiceName(serviceName), Address(address))
	c.Service = &serviceConf
	return c
}

func (c *Config) InitWeb() *Config {
	protocol := c.parser.GetString("gateway", "protocol")
	addr := c.parser.GetString("gateway", "addr")
	webConf := NewWebOptions(Protocol(protocol), Addr(addr))
	c.Web = &webConf
	return c
}
func (c *Config) InitDbType() *Config {
	c.DbType = c.parser.GetString("db_type")
	return c
}
func (c *Config) InitMysql() *Config {
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
	return c
}

func (c *Config) InitRedis() {
	addr := c.parser.GetString("redis", "addr")
	maxidle := c.parser.GetString("redis", "maxidle")
	maxactive := c.parser.GetString("redis", "maxactive")
	maxidletimeout := c.parser.GetString("redis", "maxidletimeout")
	timeout := c.parser.GetString("redis", "timeout")
	password := c.parser.GetString("redis", "password")
	dbcachepublic := c.parser.GetString("redis", "dbcachepublic")
	dbauthuser := c.parser.GetString("redis", "dbauthuser")
	dbcacheuser := c.parser.GetString("redis", "dbcacheuser")
	dbcacheadmin := c.parser.GetString("redis", "dbcacheadmin")
	dbfile := c.parser.GetString("redis", "dbfile")
	redisConf := NewRedisOptions(RedisAddr(addr), RedisMaxidle(maxidle), RedisMaxActive(maxactive),
		RedisMaxIdleTimeout(maxidletimeout), RedisTimeout(timeout), RedisPassword(password), RedisDbcachePublic(dbcachepublic),
		RedisDbauthUser(dbauthuser), RedisDbcacheUser(dbcacheuser), RedisDbcacheAdmin(dbcacheadmin), RedisDbFile(dbfile),
	)
	//fmt.Println("mysqlConf:", mysqlConf)
	c.Redis = &redisConf
}
