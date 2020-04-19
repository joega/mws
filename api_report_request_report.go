package mws

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
