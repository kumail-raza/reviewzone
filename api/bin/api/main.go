package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/gorilla/mux"

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

	var cmtService, dumpService, readerService *rpc.Client
	for i := 0; i < 3; i++ {
		select {
		case service := <-c:
			cmtService = service
		case service := <-d:
			dumpService = service
		case service := <-r:
			readerService = service
		}
	}
	defer cmtService.Close()
	defer dumpService.Close()
	defer readerService.Close()

	rt := mux.NewRouter()
	rt.HandleFunc("/comments", api.PostComments(cmtService)).Methods("POST")
	rt.HandleFunc("/comments", api.GetComments(cmtService)).Methods("GET")

	rt.HandleFunc("/dump/what", api.DumpCSV(readerService, dumpService)).Methods("GET")
	rt.HandleFunc("/csvs", api.ReadCSV(readerService)).Methods("GET")
	rt.HandleFunc("/db/csv/{csvid}", api.ReadCSVWithComments(dumpService)).Methods("GET")
	rt.HandleFunc("/db/csvs", api.ReadCSVFromDB(dumpService)).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", rt))

}
