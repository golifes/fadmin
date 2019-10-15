package config

import (
	"fadmin/tools/snowflake"
	"fmt"
	"github.com/xormplus/xorm"
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
	Node struct {
		Id int64
	}
}

var (
	engine *xorm.Engine
	Port   string
	nodeId int64
	//RedisClient *redis.Client
	//EsClient    *elastic.Client
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	config.httpServer()
	config.newNodeId()
	//config.LoadRedis()
	//config.LoadElastic()
	return
}

func NewDb() *xorm.Engine {
	fmt.Println("newDb", engine)
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
	ping := engine.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	engine.SetMaxIdleConns(c.Db.MaxIdle)
	engine.SetMaxOpenConns(c.Db.MaxOpen)
	engine.ShowSQL(c.Db.Show)
}

func (c *Config) httpServer() {
	Port = c.App.Port
}
