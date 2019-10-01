package main

import (
	"fadmin/router"
	"flag"
)

func main() {
	path := flag.String("-c", "config/config.json", "config.conf")
	router.InitRouter(*path)

}
