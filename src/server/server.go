package server

import (
	"github.com/adminibar/boatyard/src/server/middleware"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"time"
)

type Server struct {
	Addr string
}

func NewServer(c *cli.Context) *Server {
	return &Server{Addr: c.String("ui-server-addr")}
}

func (s *Server) Start() error {

	//create middleware stack
	stack := http.NotFoundHandler()
	stack = middleware.NewAssetM(stack)  //static asset middleware
	stack = middleware.NewSocketM(stack) //websocket middleware

	//create routing
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//pretent we requested index.html
		r.URL = &url.URL{Path: "index.html"}
		router.ServeHTTP(w, r)
	})

	//add router ont top of classic stack
	router.NotFoundHandler = stack

	//create HTTP server
	server := &http.Server{
		Addr:           s.Addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//start serving requests
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
