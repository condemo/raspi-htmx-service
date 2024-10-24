package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condemo/raspi-htmx-service/services/web/api/handlers"
	"github.com/condemo/raspi-htmx-service/services/web/api/middlewares"
	"github.com/condemo/raspi-htmx-service/services/web/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conf := http.NewServeMux()
	ws := http.NewServeMux()
	fs := http.FileServer(http.Dir("services/web/public/static"))
	services := http.NewServeMux()

	// Middlewares
	basicMiddleware := middlewares.MiddlewareStack(
		middlewares.Logger,
		middlewares.Recover,
	)

	// Redirect to app route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusSeeOther)
	})

	router.Handle("/app/", http.StripPrefix("/app", middlewares.RequireAuth(view)))
	router.Handle("/auth/", http.StripPrefix("/auth", basicMiddleware(auth)))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.Handle("/ws/", http.StripPrefix("/ws", ws))
	router.Handle("/conf/", http.StripPrefix("/conf", middlewares.RequireAuth(conf)))
	router.Handle("/services/", http.StripPrefix("/services", basicMiddleware(services)))

	// GRPC CONNS
	managerGrpc := newManagerGrpcClient(":8080")

	// Handlers
	authHandler := handlers.NewAuthHandler(s.store)
	viewHandler := handlers.NewViewHandler()
	wsHandler := handlers.NewWSHandler()
	confHandler := handlers.NewConfigHandler()
	servHandler := handlers.NewServiceHandler(managerGrpc)

	// Routes Load
	authHandler.RegisterRoutes(auth)
	viewHandler.RegisterRoutes(view)
	wsHandler.RegisterRoutes(ws)
	confHandler.RegisterRoutes(conf)
	servHandler.RegisterRoutes(services)

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

func newManagerGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error creating grcp manager client", err)
	}

	return conn
}
