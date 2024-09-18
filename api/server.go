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
	"github.com/condemo/raspi-htmx-service/api/handlers/middlewares"
	"github.com/condemo/raspi-htmx-service/store"
)

type ApiServer struct {
	store store.Store
	addr  string
}

func NewApiServer(addr string, s store.Store) *ApiServer {
	return &ApiServer{
		addr:  addr,
		store: s,
	}
}

func (s *ApiServer) Run() {
	router := http.NewServeMux()
	auth := http.NewServeMux()
	view := http.NewServeMux()
	ws := http.NewServeMux()
	fs := http.FileServer(http.Dir("public/static"))

	router.Handle("/app/", http.StripPrefix("/app", middlewares.RequireAuth(view)))
	router.Handle("/auth/", http.StripPrefix("/auth", middlewares.Logger(auth)))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.Handle("/ws/", http.StripPrefix("/ws", ws))

	// Handlers
	authHandler := handlers.NewAuthHandler(s.store)
	viewHandler := handlers.NewViewHandler()
	wsHandler := handlers.NewWSHandler()

	// Routes Load
	authHandler.RegisterRoutes(auth)
	viewHandler.RegisterRoutes(view)
	wsHandler.RegisterRoutes(ws)

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
