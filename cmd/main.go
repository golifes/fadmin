package main

import (
	"fadmin/router"
	"flag"
)

// @title api管理文档
// @version 1.0
// @description api文档
// @contact.name api文档
// @license.name Apache 2.0
// @host 192.168.0.5:8080

func main() {
	path := flag.String("-c", "config/config.json", "config.conf")
	router.InitRouter(*path)
}
