package pipeline

import (
	"github.com/adminibar/boatyard/src/client"
	"github.com/adminibar/boatyard/src/docker"
	"log"
)

type Pipeline struct {
	Hosts   []*docker.Host
	Clients []*client.Client
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (pl *Pipeline) AddHost(h *docker.Host) {
	log.Printf("[Pipeline] Host Added: %s", h.Addr.Host)

	//observe host and receive events
	go func() {
		c := h.Observe()
		for ev := range c {
			log.Printf("[Pipeline] Received event: %s", (<-c).Status)
			for _, client := range pl.Clients {
				go func() { client.To <- ev }()
			}
		}
	}()

	pl.Hosts = append(pl.Hosts, h)
}

func (pl *Pipeline) AddClient(client *client.Client) {
	log.Printf("[Pipeline] Client Added")

	//@todo client is ready to receive messages

	pl.Clients = append(pl.Clients, client)
}
