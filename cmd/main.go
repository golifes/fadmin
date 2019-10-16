package main

import (
	"fadmin/router"
	"flag"
)

// @title 后台管理api文档
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com
// @license.name MIT
//@host 127.0.0.1:8081

func main() {
	path := flag.String("-c", "config/config.json", "config.conf")
	router.InitRouter(*path)

}
