package server

import (
	"log"
	"net/http"
)

func New(addr string) *http.Server {
	initRoutes()
	log.Println("Server corriendo en el puerto: ", addr)
	return &http.Server{
		Addr: addr,
	}
}
