package main

import (
	"encoding/json"
	"github.com/adminibar/boatyard/src/docker"
	"github.com/adminibar/boatyard/src/server"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const DOCKER_HOST = "http://192.168.59.103:2375"

func run(c *cli.Context) {

	//need at least one docker host
	if len(c.StringSlice("host-addr")) == 0 {
		log.Fatalln("At least on docker host must be specified, use '--host-addr'")
	}

	//add to pool
	var pool []*docker.Host
	for _, addr := range c.StringSlice("host-addr") {
		purl, err := url.Parse(addr)
		if err != nil {
			log.Fatalf("Error while parsing --host-addr='%s': %s", addr, err)
		}

		log.Printf("Adding host '%s'...", purl.Host)
		host := docker.NewHost(purl)
		pool = append(pool, host)
	}

	//launch http/wss server
	server := server.NewServer(c)
	server.Start()

}

func main() {

	app := cli.NewApp()
	app.Name = "BoatYard"
	app.Usage = "real-time docker editor"
	app.Action = run

	dhost := os.Getenv("DOCKER_HOST")
	defaultHosts := &cli.StringSlice{}
	if dhost != "" {
		defaultHosts = &cli.StringSlice{dhost}
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{"ui-server-addr", "0.0.0.0:7070", "address on which to bind the ui server"},
		cli.StringSliceFlag{"host-addr", defaultHosts, "specify one ore more docker hosts, defaults to DOCKER_HOST env variable"},
	}

	app.Run(os.Args)
	return

	/**
	 * Health Check
	 */

	//create transport
	tr := &http.Transport{
		DisableCompression: true,
	}

	//create client
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get(DOCKER_HOST + "/_ping")
	if err != nil {
		log.Fatalln(err)
	}

	//check health
	log.Printf("Host status is: '%s'", resp.Status)

	/**
	 * Event Hook
	 */

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
	resp, err = cconn.Do(req)
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
