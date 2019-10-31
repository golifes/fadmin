package config

import (
	"fadmin/tools/snowflake"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
	"log"
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
	Es struct {
		Host  string //es host
		Index string // es index
	}

	App struct {
		Port string
	}
	Node struct {
		Id int64
	}
}

var (
	engine      *xorm.Engine
	Port        string
	RedisClient *redis.Client
	EsClient    *elastic.Client
	EsIndex     string
	//RedisClient *redis.Client
	//EsClient    *elastic.Client
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	config.httpServer()
	_, _ = config.newNodeId()
	config.LoadRedis()
	config.LoadElastic()
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

func (c *Config) LoadElastic() {
	var err error
	EsIndex = c.Es.Index
	//EsClient, err = elastic.NewSimpleClient(elastic.SetURL(c.Es.Host))
	//if err != nil {
	//	log.Printf("elastic conn is error %s", err)
	//}
	//
	//info := elasticsearch.Config{
	//	Addresses: []string{""},
	//	Transport: &http.Transport{
	//		MaxIdleConnsPerHost:   10,
	//		ResponseHeaderTimeout: 500 * time.Millisecond,
	//		DialContext:           (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
	//		TLSClientConfig: &tls.Config{
	//			MinVersion: tls.VersionTLS11,
	//			// ...
	//		},
	//	},
	//}
	if EsClient, err = elastic.NewClient(elastic.SetURL(c.Es.Host)); err != nil {
		log.Printf("elastic conn is error %v", err)
	}
}
