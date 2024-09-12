package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/condemo/raspi-htmx-service/api"
	"github.com/condemo/raspi-htmx-service/store"
)

func main() {
	addr := flag.String("p", ":4000", "service port")
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
