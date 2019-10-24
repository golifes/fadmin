package config

import (
	"fadmin/tools/snowflake"
	"fmt"
	"github.com/xormplus/xorm"
	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v7"
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

	Redis struct {
		Dns      string
		Pwd      string
		PoolSize int
		MaxIdle  int
		Db       int
	}

	App struct {
		Port string
	}
	Node struct {
		Id int64
	}
}

var (
	engine *xorm.Engine
	Port   string
	//RedisClient *redis.Client
	//EsClient    *elastic.Client
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	config.httpServer()
	_, _ = config.newNodeId()
	//config.LoadRedis()
	//config.LoadElastic()
	return
}

func NewDb() *xorm.Engine {
	return engine
}

func (c *Config) newNodeId() (*snowflake.Node, error) {
	node, err := snowflake.NewNode(c.Node.Id)
	return node, err
}

func NewNodeId() int64 {
	config := Config{}
	node, err := config.newNodeId()
	if err != nil {
		return 0
	}
	return node.Id()
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

	engine, err = xorm.NewEngine("mysql", dns)

	if err != nil || engine.Ping() != nil {
		panic(err)
	}

	engine.SetMaxIdleConns(c.Db.MaxIdle)
	engine.SetMaxOpenConns(c.Db.MaxOpen)
	engine.ShowSQL(c.Db.Show)

}

func (c *Config) httpServer() {
	Port = c.App.Port
}

func (c *Config) LoadRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Dns,
		Password: c.Redis.Pwd, // no password set
		DB:       c.Redis.Db,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}
