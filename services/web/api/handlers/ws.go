package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/a-h/templ"
	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/components"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type WSHandler struct {
	sysInfoConn pb.SysInfoServiceClient
	mu          *sync.RWMutex
	conns       map[*websocket.Conn]struct{}
}

func NewWSHandler(siConn *grpc.ClientConn) *WSHandler {
	si := pb.NewSysInfoServiceClient(siConn)
	return &WSHandler{
		sysInfoConn: si,
		mu:          new(sync.RWMutex),
		conns:       make(map[*websocket.Conn]struct{}),
	}
}

func (h *WSHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/info", MakeHandler(h.getConn))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *WSHandler) getConn(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	h.handleWS(conn)
	return nil
}

func (h *WSHandler) handleWS(c *websocket.Conn) {
	fmt.Println("New Connection:", c.RemoteAddr())

	h.mu.Lock()
	h.conns[c] = struct{}{}
	h.mu.Unlock()

	s := make(chan struct{})

	go h.writeLoop(c, s)
	go h.readLoop(c, s)
}

func (h *WSHandler) writeLoop(c *websocket.Conn, s chan struct{}) {
	t := time.NewTicker(config.UsConf.InfoConf.InfoTick)
	for {
		select {
		case <-t.C:
			si, err := h.sysInfoConn.GetInfo(context.Background(), &pb.GetInfoRequest{})
			if err != nil {
				log.Fatalf("something wrong with GetInfo %v \n", err)
			}

			tmpl, err := templ.ToGoHTML(context.Background(), components.Infobar(si.GetSisInfo()))
			if err != nil {
				fmt.Println("error converting Infobar to html:", err)
				return
			}

			uptimeHTML, err := templ.ToGoHTML(context.Background(), components.UptimeLabel(si.GetSisInfo().GetUptime()))
			if err != nil {
				fmt.Println("error converting UptimeLabel to html:", err)
				return
			}

			c.WriteMessage(websocket.TextMessage, []byte(tmpl))
			c.WriteMessage(websocket.TextMessage, []byte(uptimeHTML))

		case <-s:
			h.mu.Lock()
			delete(h.conns, c)
			h.mu.Unlock()
			fmt.Printf("Connection with %s closed\n", c.RemoteAddr())
			return
		}
	}
}

func (h *WSHandler) readLoop(c *websocket.Conn, s chan struct{}) {
	for {
		if _, _, err := c.NextReader(); err != nil {
			c.Close()
			close(s)
			break
		}
	}
}
