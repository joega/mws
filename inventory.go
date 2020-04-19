package mws

//FulfillmentInventoryService 库存管理
type FulfillmentInventoryService struct {
	*Client
}

//FulfillmentInventory FulfillmentInventory
func FulfillmentInventory() *FulfillmentInventoryService {
	return &FulfillmentInventoryService{
		Client: newClient("/FulfillmentInventory/2010-10-01", "2010-10-01"),
	}
}

//InventorySupplyResult InventorySupplyResult
type InventorySupplyResult struct {
	NextToken           string
	InventorySupplyList []*InventorySupplyList `xml:"InventorySupplyList>member"`
}

//InventorySupplyList InventorySupplyList
type InventorySupplyList struct {
	Condition             string
	TotalSupplyQuantity   int
	InStockSupplyQuantity int
	FNSKU                 string
	ASIN                  string
	SellerSKU             string
}

// ListInventorySupply operation returns information about the availability of
// inventory that a seller has in Amazon's fulfillment network and in current inbound
// shipments. You can check the current availability status for your Fulfillment by
// Amazon inventory as well as discover when availability status changes.
func (s *FulfillmentInventoryService) ListInventorySupply(c *Credential, params ...Values) (string, *InventorySupplyResult, error) {
	data := ActionValues("ListInventorySupply")
	data.SetAll(params...)

	var response struct {
		BaseResponse
		InventorySupplyResult *InventorySupplyResult `xml:"ListInventorySupplyResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.InventorySupplyResult, nil
}

// ListInventorySupplyByNextToken operation returns the next page of information
// about the availability of a seller's inventory using the NextToken value that
// was returned by your previous request to either ListInventorySupply or
// ListInventorySupplyByNextToken. If NextToken is not returned, there are no more
// pages to return.
func (s *FulfillmentInventoryService) ListInventorySupplyByNextToken(c *Credential, nextToken string) (string, *InventorySupplyResult, error) {
	data := ActionValues("ListInventorySupply")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		InventorySupplyResult *InventorySupplyResult `xml:"ListInventorySupplyByNextTokenResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.InventorySupplyResult, nil
}
