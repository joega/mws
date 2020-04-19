package mws
import "time"

//ReportService .
type ReportService struct {
	*Client
}

//ReportListResult .
type ReportListResult struct {
	HasNext   bool
	NextToken string
	Reports   []*ReportInfo `xml:"ReportInfo"`
}

//ReportInfo .
type ReportInfo struct {
	ReportID        string `xml:"ReportId"`
	ReportRequestID string `xml:"ReportRequestId"`
	ReportType      string
	Acknowledged    bool
	AvailableDate   *time.Time
}

//ReportRequestListResult .
type ReportRequestListResult struct {
	NextToken         string
	HasNext           bool
	ReportRequestInfo []*ReportRequestInfo
}

//ReportRequestInfo .
type ReportRequestInfo struct {
	ReportRequestID        string `xml:"ReportRequestId"`
	ReportType             string
	StartDate              time.Time
	EndDate                time.Time
	Scheduled              bool
	SubmittedDate          time.Time
	ReportProcessingStatus string
	GeneratedReportID      string `xml:"GeneratedReportId"`
	StartedProcessingDate  time.Time
	CompletedDate          time.Time
}

//Reports .
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


// GetReportList operation returns a list of reports that were created in the
// previous 90 days that match the query parameters. A maximum of 100 results
// can be returned in one request. If there are additional results to return,
// HasNext is returned set to true in the response. To retrieve all the results,
// you can pass the value of the NextToken parameter to the GetReportListByNextToken
// operation iteratively until HasNext is returned set to false.
func (s *ReportService) GetReportList(c *Credential, params ...Values) (string, *ReportListResult, error) {
	data := ActionValues("GetReportList")
	data.SetAll(params...)
	// data.SetInt("MaxCount", maxCount)
	// data.SetTime("AvailableFromDate", availableFromDate)
	// data.SetTime("AvailableToDate", availableToDate)
	// data.SetBool("Acknowledged", acknowledged)
	// data.Sets("ReportTypeList.Type", reportTypeList...)
	// data.Sets("ReportRequestIdList.Id", reportRequestIDList...)

	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}

// GetReportListByNextToken operation returns a list of reports that match the
// query parameters, using the NextToken, which was supplied by a previous call
// to either GetReportListByNextToken or a call to GetReportList, where the value
// of HasNext was true in the previous call.
func (s *ReportService) GetReportListByNextToken(c *Credential, nextToken string) (string, *ReportListResult, error) {
	data := ActionValues("GetReportListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListByNextTokenResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}

// GetReportRequestList operation returns a list of report requests that match the
// query parameters. You can specify query parameters for report status, date range,
// and report type. The list contains the ReportRequestId for each report request.
// You can obtain ReportId values by passing the ReportRequestId values to the
// GetReportList operation. In the first request, a maximum of 100 report requests
// are returned. If there are additional report requests to return, HasNext is returned
// set to true in the response . To retrieve all the results, you can pass the value of
// the NextToken parameter to call GetReportRequestListByNextToken operation iteratively
// until HasNext is returned set to false.
func (s *ReportService) GetReportRequestList(c *Credential, params ...Values) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestList")
	data.SetAll(params...)
	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestListResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}

// GetReportRequestListByNextToken operation returns a list of report requests
// that match the query parameters. This operation uses the NextToken, which was
// supplied by a previous request to either GetReportRequestListByNextToken or a
// request to GetReportRequestList, where the value of HasNext was true in that
// previous request.
func (s *ReportService) GetReportRequestListByNextToken(c *Credential, nextToken string) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestLisByNextTokentResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}

// RequestReport operation creates a report request. Amazon MWS processes the report
// request and when the report is completed, sets the status of the report request to
// _DONE_. Reports are retained for 90 days.  You specify what marketplaces you want a
// report to cover by supplying a list of marketplace IDs to the optional MarketplaceIdList
// request parameter when you call the RequestReport operation. If you do not specify a
// marketplace ID, your home marketplace ID is used. Note that the MarketplaceIdList
// request parameter is not used in the Japan marketplace. The RequestReport operation has a
// maximum request quota of 15 and a restore rate of one request every minute.
func (s *ReportService) RequestReport(c *Credential, reportType string, params ...Values) (string, *ReportRequestInfo, error) {
	data := ActionValues("RequestReport")
	data.Set("ReportType", reportType)
	data.SetAll(params...)
	// data.Set("ReportOptions", reportOptions)
	// data.SetTime("StartDate", startDate)
	// data.SetTime("EndDate", endDate)
	// data.Sets("MarketplaceIdList.Id", marketplaceIDList...)

	var response struct {
		BaseResponse
		ReportRequestInfo *ReportRequestInfo `xml:"RequestReportResult>ReportRequestInfo"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestInfo, nil
}


