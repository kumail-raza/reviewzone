package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/rpc"

	boom "github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/reviewzone/comments"
)

//TestRoute TestRoute
func TestRoute(readerService, dumpService, cmtService *rpc.Client) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(csvs[0][0]))
	}
}

func Respond(w http.ResponseWriter, i interface{}) {
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(i)
	if err != nil {
		boom.Internal(w, err)
		return
	}

}
