package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/store"
	"github.com/condemo/raspi-htmx-service/services/web/api"
)

func main() {
	addr := flag.String("p", config.ServicesConfig.WebServPort, "service port")
	flag.Parse()

	sqlStorage := store.NewPostgresStore()
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(db)

	apiServer := api.NewApiServer(*addr, store)
	fmt.Println("Server Running on port", *addr)
	apiServer.Run()
}
