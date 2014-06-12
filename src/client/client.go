package client

import (
	"github.com/adminibar/boatyard/src/docker"
	"github.com/gorilla/websocket"
	"log"
	// "time"
)

type Client struct {
	Conn *websocket.Conn
	To   chan docker.Event
}

func NewClient(c *websocket.Conn) *Client {
	return &Client{
		Conn: c,
	}
}

//open the connection forever
func (c *Client) Open() {
	defer c.Conn.Close()

	//@todo replace with some
	//logic that determines how the connection
	//is kept open
	for to := range c.To {
		log.Println("Message to Client:", to)
		// time.Sleep(time.Second)

		// //write message, if we error assume the client left, leave loop
		// err := c.Conn.WriteMessage(websocket.TextMessage, []byte("hi"))
		// if err != nil {
		// 	log.Println(err)
		// 	break
		// }
	}

}
