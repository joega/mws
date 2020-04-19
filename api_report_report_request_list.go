package mws

import "time"

//ReportRequestListResult 获取报告请求列表结果
type ReportRequestListResult struct {
	NextToken         string
	HasNext           bool
	ReportRequestInfo []*ReportRequestInfo
}

//ReportRequestInfo 获取报告请求列表结果
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
