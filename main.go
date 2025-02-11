package main

import (
	"fmt"

	"github.com/danielnicolae28/webserver/api"
	"github.com/danielnicolae28/webserver/server"
)

func main() {
	fmt.Println("Web server in golang")
	api.DataApi()
	server.Webserver()
}
