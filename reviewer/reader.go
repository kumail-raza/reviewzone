package reviewer

import (
	"encoding/csv"
	"os"
)

//ReadCSVFile ReadCSVFile
func (r *Reviewer) readCSVFile(path string) ([][]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	fLine, err := reader.ReadAll()

	return fLine, nil
}
