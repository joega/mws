package mws

//ReportService 报表服务
type ReportService struct {
	*Client
}

//Reports 创建报表服务
func Reports() *ReportService {
	return &ReportService{
		Client: newClient("/", "2009-01-01"),
	}
}

// The GetReport operation returns the contents of a report and the Content-MD5 header
// for the returned report body.
func (s *ReportService) GetReport(c *Credential, ReportID string) ([]byte, error) {
	data := ActionValues("GetReport")
	data.Set("ReportId", ReportID)
	return s.GetBytes(c, data)
}
