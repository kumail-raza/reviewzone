package api

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/minhajuddinkhan/reviewzone/comments"
)

//TestRoute TestRoute
func TestRoute(readerService, dumpService, cmtService *rpc.Client) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var csvs [][]string
		err := readerService.Call("Service.ReadCSVFile", "/tmp/text.csv", &csvs)
		if err != nil {
			panic(err)
		}

		spew.Dump(csvs)
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
		w.Write([]byte(csvs[0][0]))
		w.WriteHeader(http.StatusOK)
	}
}
