package main

import (
	"CatRegistry/src/internal/router"
	"fmt"
	"net/http"
	"os"
)

var (
	uri      string
	port     string
	hostname string
	hostport string
)

func init() {
	uri = os.Getenv("MONGODB_URI")
	port = os.Getenv("MONGODB_PORT")
	hostname = os.Getenv("HOSTNAME")
	hostport = os.Getenv("HOSTPORT")
}

func main() {
	doServering()
}

func doServering() {
	if hostname == "" {
		hostname = "127.0.0.1"
	}
	if hostport == "" {
		hostport = "8080"
	}
	serv := router.CreateNewHTTPServer()
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", hostname, hostport),
		Handler: serv,
	}
	fmt.Println("Serving http")
	srv.ListenAndServe()
}
