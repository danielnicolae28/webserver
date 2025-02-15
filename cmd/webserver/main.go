package main

import (
	"fmt"

	api "github.com/danielnicolae28/webserver/api/server"
)

func main() {
	fmt.Println("Web server in golang")
	api.Webserver()
}
