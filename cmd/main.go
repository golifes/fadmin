package main

import (
	"fadmin/cmd/docs"
	"fadmin/router"
	"flag"
)

func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	path := flag.String("-c", "config/config.json", "config.conf")
	router.InitRouter(*path)

}
