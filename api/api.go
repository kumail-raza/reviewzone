package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/minhajuddinkhan/reviewzone/admin"
)

func main() {

	// r := reviewer.Reviewer{}
	// csvs, err := r.ReadCSVFile("text.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// d := dump.Dumper{}
	// ids, err := d.DumpCSV(csvs)
	// spew.Dump(ids)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	a := admin.Admin{}
	x, _ := a.GetCSVs()
	// x := comments.Comments{}
	// x.Comments = []string{"Nice one"}
	// err = x.AddOnCSV("5bd372922cbccbf692e7423a")
	// if err != nil {
	// 	panic(err)
	// }
	spew.Dump(x)
}
