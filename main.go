package main

import (
	"github.com/adminibar/boatyard/src/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

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
		Addr:           "127.0.0.1:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//start serving requests
	log.Printf("HTTP Listening on '%s'...", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
