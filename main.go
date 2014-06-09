package main

import (
	"github.com/adminibar/boatyard/src/docker"
	"github.com/adminibar/boatyard/src/server"
	"github.com/codegangsta/cli"
	"log"
	"net/url"
	"os"
	"time"
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

		log.Printf("Connecting to host '%s'...", purl.Host)
		host, err := docker.NewHost(purl)
		if err != nil {
			log.Fatalf("Error setting up host --host-addr='%s': %s", addr, err)
		}

		//wait until healthy
		healthy := make(chan bool)
		go func() {
			for {

				//check health
				status, err := host.Ping()
				if err != nil {
					log.Printf("Host[%s] health check failed: '%s'", host.Addr, err)
				}

				//if where only indicate healty
				if status == "200 OK" {
					healthy <- true
					break
				}

				dur := time.Second * 5
				log.Printf("Host[%s] is not online retrying in (%s)", purl.Host, dur)
				time.Sleep(dur)
			}
		}()

		//block until host is up
		<-healthy
		log.Printf("Host[%s] is healthy!", purl.Host)

		//@todo proceed with connecting
		host.Observe()
		pool = append(pool, host)

	}

	//launch http/wss server
	server := server.NewServer(c)
	log.Printf("Starting HTTP server on '%s'...", server.Addr)
	err := server.Start()
	if err != nil {
		log.Fatalf("Error while starting HTTP server '%s': '%s'", server.Addr, err)
	}

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
}
