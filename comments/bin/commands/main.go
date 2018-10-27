package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/minhajuddinkhan/reviewzone/comments"
)

func main() {

	service := new(comments.Service)
	rpc.Register(service)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(l, nil)
}
