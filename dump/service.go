package dump

//Service This is a dumping micro service
type Service struct{}

//DumpCSV DumpCSV
func (ds *Service) DumpCSV(in [][]string, out *[]string) error {

	d := Dumper{}
	csv, err := d.dumpCSV(in)
	if err != nil {
		return err
	}
	*out = csv
	return nil
}
