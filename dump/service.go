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

//ReadCSVFromDB ReadCSVFromDB
func (ds *Service) ReadCSVFromDB(in string, out *[]Format) error {
	d := Dumper{}
	csvs, err := d.readCSVFromDB()
	if err != nil {
		return err
	}
	*out = csvs
	return nil
}

func (ds *Service) ReadOneCSV(in string, out *Format) error {

	d := Dumper{}
	csv, err := d.readOneCSV(in)
	if err != nil {
		return err
	}
	*out = csv
	return nil
}
