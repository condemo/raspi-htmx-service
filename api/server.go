package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condemo/raspi-htmx-service/api/handlers"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() {
	router := http.NewServeMux()
	view := http.NewServeMux()
	fs := http.FileServer(http.Dir("public/static"))

	router.Handle("/", view)
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handlers
	viewHandler := handlers.NewViewHandler()

	// Routes Load
	viewHandler.RegisterRoutes(view)

	server := http.Server{
		Addr:         s.addr,
		Handler:      router,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 3,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// server.Shutdown ends the execution of the program
	// after waiting for all active connections to finish or 30 seconds to pass
	server.Shutdown(ctx)
}
