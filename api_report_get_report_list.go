package mws

import "time"

//ReportListResult 获取报告列表的结果模型
type ReportListResult struct {
	HasNext   bool
	NextToken string
	Reports   []*ReportInfo `xml:"ReportInfo"`
}

//ReportInfo 报告信息在结果列表中的模型
type ReportInfo struct {
	ReportID        string `xml:"ReportId"`
	ReportRequestID string `xml:"ReportRequestId"`
	ReportType      string
	Acknowledged    bool
	AvailableDate   *time.Time
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
