package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/minhajuddinkhan/reviewzone/comments"
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

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		var csvs [][]string
		err := readerService.Call("Service.ReadCSVFile", "text.csv", &csvs)
		if err != nil {
			panic(err)
		}

		var csvIDs []string
		dumpService.Call("Service.DumpCSV", csvs, &csvIDs)

		var commentIds []string
		req := comments.AddCommentRequest{
			Comments: []string{"Nice"},
			CsvID:    csvIDs[0],
		}
		err = cmtService.Call("Service.OnCsvComment", req, &commentIds)
		if err != nil {
			log.Fatal(err)
		}

		var comments []comments.Comment
		cmtService.Call("Service.GetComments", csvIDs[0], &comments)
		spew.Dump(comments)
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":6000", nil)

}
