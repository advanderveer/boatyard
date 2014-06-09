package middleware

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

/**
 * Websocket Middle ware that attempts to
 * return embedded web resources
 */
type SocketM struct {
	next     http.Handler
	upgrader websocket.Upgrader
}

//create the middleware
func NewSocketM(handler http.Handler) *SocketM {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &SocketM{next: handler, upgrader: upgrader}
}

//interface implementation for middleware
func (s *SocketM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/ws" {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		//close connection as soon as the loop ends
		defer conn.Close()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)

			//write message, if we error assume the client left, leave loop
			err := conn.WriteMessage(websocket.TextMessage, []byte("hi"))
			if err != nil {
				log.Println(err)
				break
			}
		}

	} else {
		s.next.ServeHTTP(w, r)
	}
}
