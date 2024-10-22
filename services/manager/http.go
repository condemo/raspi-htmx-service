package manager

import (
	"log"
	"net/http"

	handlers "github.com/condemo/raspi-htmx-service/services/manager/handlers/manager"
	"github.com/condemo/raspi-htmx-service/services/manager/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() {
	r := http.NewServeMux()

	managerService := service.NewManagerService()
	managerHandler := handlers.NewManagerHttpHandler(managerService)
	managerHandler.RegisterRoutes(r)

	log.Println("HyperMedia http on port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, r))
}
