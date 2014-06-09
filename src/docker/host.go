package docker

import (
	"net/url"
)

type Host struct {
	Addr *url.URL
}

func NewHost(addr *url.URL) *Host {
	return &Host{Addr: addr}
}
