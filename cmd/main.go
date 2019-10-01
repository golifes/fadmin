package main

import (
	"fadmin/router"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	router.InitRouter("config/config.json")

}
