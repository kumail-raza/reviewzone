package api

import (
	"encoding/json"
	"net/http"
	"net/rpc"

	boom "github.com/darahayes/go-boom"
	"github.com/gorilla/mux"

	"github.com/minhajuddinkhan/reviewzone/comments"
)

//CommentRequest CommentRequest
type CommentRequest struct {
	Comments []string `json:"comments"`
	CsvID    string   `json:"csvid"`
}

//GetComments GetComments
func GetComments(cmtService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		csvID := mux.Vars(r)["csv"]
		if csvID == "" {
			boom.NotFound(w, "csv id not found.")
			return
		}
		var comments []comments.Comment
		err := cmtService.Call("Service.GetComments", csvID, &comments)
		if err != nil {
			boom.BadGateway(w, err.Error())
		}

		Respond(w, comments)
		return
	}
}

//PostComments PostComments
func PostComments(cmtService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var cr CommentRequest
		err := decoder.Decode(&cr)
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}
		var commentIds []string
		err = cmtService.Call("Service.OnCsvComment", comments.AddCommentRequest{Comments: cr.Comments, CsvID: cr.CsvID}, &commentIds)
		if err != nil {
			boom.BadData(w, err.Error())
		}

		Respond(w, commentIds)
		return

	}
}
