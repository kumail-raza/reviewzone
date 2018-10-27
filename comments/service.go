package comments

//AddCommentRequest AddCommentRequest
type AddCommentRequest struct {
	CsvID    string
	Comments []string
}

//Service Service
type Service struct {
}

//GetComments GetComments
func (cs *Service) GetComments(in string, out *[]Comment) error {

	c := Comment{}
	comments, err := c.getComments(in)
	if err != nil {
		return err
	}
	*out = comments
	return nil
}

//OnCsvComment OnCsvComment
func (cs *Service) OnCsvComment(in AddCommentRequest, out *[]string) error {

	c := Comment{}
	csvIds, err := c.addOnCSV(in.CsvID, in.Comments...)
	if err != nil {
		return err
	}
	*out = csvIds
	return nil
}
