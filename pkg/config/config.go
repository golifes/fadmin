package config

import (
	"fmt"
	//"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golifes/sqlo"
	//"github.com/olivere/elastic/v7"
	//"github.com/xormplus/xorm"
)

type Config struct {
	Db struct {
		Host    string
		User    string
		Pwd     string
		Db      string
		Show    bool
		Port    int
		MaxOpen int
		MaxIdle int
	}
	App struct {
		Port string
	}
}

var (
	engine sqlo.Engine
	Port   string
	//RedisClient *redis.Client
	//EsClient    *elastic.Client
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	config.httpServer()
	//config.LoadRedis()
	//config.LoadElastic()
	return
}

func NewDb() sqlo.Engine {
	fmt.Println("newDb", engine)
	return engine
}

func NewHttpPort() string {
	return Port
}
func (c *Config) loadDb() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Db.User,
		c.Db.Pwd,
		c.Db.Host,
		c.Db.Db)
	fmt.Println(dns)

	engine, err = sqlo.Connect(dns)
	ping := engine.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	engine.DB().SetMaxIdleConns(c.Db.MaxIdle)
	engine.DB().SetMaxOpenConns(c.Db.MaxOpen)
}

func (c *Config) httpServer() {
	Port = c.App.Port
}
