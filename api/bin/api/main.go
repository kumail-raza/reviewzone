package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/minhajuddinkhan/reviewzone/api"
)

const (
	//CommentsAddr CommentsAddr
	CommentsAddr = "comments:3000"
	//DumpAddr DumpAddr
	DumpAddr = "dump:4000"
	//ReaderAddr ReaderAddr
	ReaderAddr = "reviewer:5000"
)

func getService(addr string) *rpc.Client {
	s, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func main() {

	c := make(chan *rpc.Client)
	d := make(chan *rpc.Client)
	r := make(chan *rpc.Client)
	go func() { c <- getService(CommentsAddr) }()
	go func() { d <- getService(DumpAddr) }()
	go func() { r <- getService(ReaderAddr) }()

	cmtService := <-c
	dumpService := <-d
	readerService := <-r

	defer dumpService.Close()
	defer cmtService.Close()
	defer readerService.Close()

	http.HandleFunc("/go", api.TestRoute(readerService, dumpService, cmtService))

	http.ListenAndServe(":6000", nil)

}
