package main

import (
	"flag"
	"fmt"

	"github.com/condemo/raspi-htmx-service/api"
)

func main() {
	addr := flag.String("p", ":4000", "service port")
	flag.Parse()

	apiServer := api.NewApiServer(*addr)
	fmt.Println("Server Running on port", *addr)
	apiServer.Run()
}
