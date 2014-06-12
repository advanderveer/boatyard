package middleware

import (
	"github.com/adminibar/boatyard/src/client"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/**
 * Websocket Middle ware that attempts to
 * return embedded web resources
 */
type SocketM struct {
	next     http.Handler
	upgrader websocket.Upgrader
	Clients  chan *client.Client
}

//create the middleware
func NewSocketM(handler http.Handler) *SocketM {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &SocketM{
		next:     handler,
		upgrader: upgrader,
		Clients:  make(chan *client.Client),
	}
}

//interface implementation for middleware
func (s *SocketM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/ws" {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		//new client by socket, blocks until someone handles the client
		//from this point on the connection is hijacked
		s.Clients <- client.NewClient(conn)
	} else {
		s.next.ServeHTTP(w, r)
	}
}
