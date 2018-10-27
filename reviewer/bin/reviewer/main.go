package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/minhajuddinkhan/reviewzone/reviewer"
)

func main() {

	service := new(reviewer.Service)
	rpc.Register(service)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(l, nil)
}
