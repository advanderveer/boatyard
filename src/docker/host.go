package docker

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"encoding/json"
	"io"
	"log"
	"net/http/httputil"
)

type Host struct {
	Addr  *url.URL
	Agent *http.Client
}

type Event struct {
	Status string
	ID     string
	From   string
	Time   int64
}

// Create a new host based on a given url
func NewHost(addr *url.URL) (*Host, error) {

	//create custom agent Dial
	adial, err := net.DialTimeout("tcp", addr.Host, time.Second*5)
	if err != nil {
		return nil, err
	}

	//create transport for agent
	atr := &http.Transport{
		DisableCompression: true,
		Dial:               func(string, string) (net.Conn, error) { return adial, nil },
	}

	//create http client
	agent := &http.Client{
		Transport: atr,
	}

	//default to supported http scheme
	if addr.Scheme == "" || addr.Scheme == "tcp" {
		addr.Scheme = "http"
	}

	return &Host{
		Addr:  addr,
		Agent: agent,
	}, nil
}

// Ping theh docker host and see if its healthy
func (h *Host) Ping() (string, error) {
	resp, err := h.Agent.Get(h.Addr.String() + "/_ping")
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

//
// @todo impelment Close() (stop observing)
//

// Start observing the hosts event endpoint
// Generator pattern, returns a channel of api events
func (h *Host) Observe() <-chan Event {

	//@TODO Refactor this function

	//1. create connection
	conn, err := net.Dial("tcp", h.Addr.Host)
	if err != nil {
		log.Fatalln(err)
	}

	//2. create client for connection
	cconn := httputil.NewClientConn(conn, nil)

	//3. create request
	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//4. execute request
	resp, err := cconn.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//start sending event across channel
	c := make(chan Event)
	go func(resp *http.Response, conn *httputil.ClientConn, c chan Event) {
		defer conn.Close()
		defer resp.Body.Close()

		//keep decoding the body
		decoder := json.NewDecoder(resp.Body)
		for {
			var event Event
			err := decoder.Decode(&event)
			if err != nil {

				//close down connection when EOF is received
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					log.Println("Host[%s] sent EOF while observing", h.Addr.Host)
					break
				} else {
					//@todo send to error channel
					log.Fatalln(err)
				}

			} else {
				//log and send to channel, block until received
				log.Printf("Host[%s] sent event: %s", h.Addr.Host, event.Status)
				c <- event
			}
		}

		//@todo atempt to reconnect
		log.Println("/event connection closed")

	}(resp, cconn, c)

	return c
}
