package main

import (
	"fmt"

	"github.com/danielnicolae28/webserver/api/handler"
	"github.com/danielnicolae28/webserver/api/server"
)

func main() {
	fmt.Println("Web server in golang")
	api.DataApi()
	server.Webserver()
}
