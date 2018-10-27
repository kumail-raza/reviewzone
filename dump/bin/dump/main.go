package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/minhajuddinkhan/reviewzone/dump"
)

func main() {

	service := new(dump.Service)
	rpc.Register(service)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(l, nil)
}
