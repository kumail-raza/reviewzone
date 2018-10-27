package reviewer

//Service Reviewer Service
type Service struct{}

//ReadCSVFile ReadCSVFile
func (rs *Service) ReadCSVFile(in string, out *[][]string) error {

	r := Reviewer{}
	csvs, err := r.readCSVFile(in)
	if err != nil {
		return err
	}
	*out = csvs
	return nil
}
