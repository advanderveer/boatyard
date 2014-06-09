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

// Start observing the hosts event endpoint
func (h *Host) Observe() {

	//1. create connection
	conn, err := net.Dial("tcp", "192.168.59.103:2375")
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

	type Event struct {
		Status string
		ID     string
		From   string
		Time   int64
	}

	//listen to events
	go func(res *http.Response, conn *httputil.ClientConn) {
		defer conn.Close()
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)

		for {
			var event Event
			err := decoder.Decode(&event)
			if err != nil {

				//close down connection when EOF is received
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					log.Println("Received EOF from /event entpoint")
					break
				} else {
					//@todo send to error channel
					log.Fatalln(err)
				}

			} else {

				//@todo send to event channel
				log.Println(event)
			}
		}

		//@todo atempt to reconnect
		log.Println("/event connection closed")

	}(resp, cconn)

}
