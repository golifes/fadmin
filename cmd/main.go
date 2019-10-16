package main

import (
	"fadmin/cmd/docs"
	"fadmin/router"
	"flag"
)

// @title Swagger Example API12222
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	path := flag.String("-c", "config/config.json", "config.conf")
	router.InitRouter(*path)

}
