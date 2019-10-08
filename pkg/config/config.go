package config

import (
	"database/sql"
	"fmt"
	//"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
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
	Db   *sql.DB
	Port string
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

func NewDb() *sql.DB {
	return Db
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
	Db, err = sql.Open("mysql", dns)

	ping := Db.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	Db.SetMaxIdleConns(c.Db.MaxIdle)
	Db.SetMaxOpenConns(c.Db.MaxOpen)
}

func (c *Config) httpServer() {
	Port = c.App.Port
}
