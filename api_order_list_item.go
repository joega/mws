package mws

//OrderItemsResult .
type OrderItemsResult struct {
	NextToken     string
	AmazonOrderID string       `xml:"AmazonOrderId"`
	OrderItems    []*OrderItem `xml:"OrderItems>OrderItem"`
}

//OrderItem  订单商品
type OrderItem struct {
	ASIN                          string      `xml:"ASIN"`                           //商品的亚马逊标准识别号 (ASIN)。
	SellerSKU                     string      `xml:"SellerSKU"`                      //商品的卖家 SKU。
	OrderItemID                   string      `xml:"OrderItemId"`                    //亚马逊定义的订单商品识别号。
	Title                         string      `xml:"Title"`                          //商品名称
	NumberOfItems                 int         `xml:"ProductInfo>NumberOfItems"`      //一件，包含是数量
	QuantityOrdered               int         `xml:"QuantityOrdered"`                //下单的商品数量
	QuantityShipped               int         `xml:"QuantityShipped"`                //已配送的商品数量。
	ItemPrice                     Money       `xml:"ItemPrice"`                      //订单商品的总价。注意，订单商品指的商品和数量。这意味着，ItemPrice 的价值等于商品售价乘以订购数量。请注意： ItemPrice 不包括ShippingPrice 和GiftWrapPrice
	ItemTax                       Money       `xml:"ItemTax"`                        //商品价格的税费。
	ShippingPrice                 Money       `xml:"ShippingPrice"`                  //运费
	GiftWrapPrice                 Money       `xml:"GiftWrapPrice"`                  //商品的礼品包装金额。
	ShippingTax                   Money       `xml:"ShippingTax"`                    //运费的税费。
	GiftWrapTax                   Money       `xml:"GiftWrapTax"`                    //礼品包装金额的税费。
	ShippingDiscount              Money       `xml:"ShippingDiscount"`               //运费的折扣。
	PromotionDiscount             Money       `xml:"PromotionDiscount"`              //报价中的全部促销折扣总计。
	PromotionDiscountTax          Money       `xml:"PromotionDiscountTax"`           //报价中的全部促销折扣总计。
	PromotionIDs                  []string    `xml:"PromotionIds"`                   //PromotionId 元素列表。
	CODFee                        Money       `xml:"CODFee"`                         //COD 服务费用。 注： CODFee 是仅在日本 (JP) 使用的响应元素。
	CODFeeDiscount                Money       `xml:"CODFeeDiscount"`                 //货到付款费用的折扣。注： CODFeeDiscount 是仅在日本 (JP) 使用的响应元素。
	IsGift                        bool        `xml:"IsGift"`                         //买家提供的礼品消息。
	GiftMessageText               string      `xml:"GiftMessageText"`                //买家提供的礼品消息。
	GiftWrapLevel                 string      `xml:"GiftWrapLevel"`                  //买家指定的礼品包装等级。
	InvoiceData                   InvoiceData `xml:"InvoiceData"`                    //发票信息（仅适用于中国）。注： InvoiceData 响应元素仅适用于中国 (CN)。
	ConditionNote                 string      `xml:"ConditionNote"`                  //卖家描述的商品状况。
	ConditionID                   string      `xml:"ConditionId"`                    //商品的状况。New, Used, Collectible, Refurbished, Preorder, Club
	ConditionSubtypeID            string      `xml:"ConditionSubtypeId"`             //商品的子状况。New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, Any, Other
	ScheduledDeliveryStartDate    string      `xml:"ScheduledDeliveryStartDate"`     //订单预约送货上门的开始日期（目的地时区）。日期格式为 ISO 8601。
	ScheduledDeliveryEndDate      string      `xml:"ScheduledDeliveryEndDate"`       //订单预约送货上门的终止日期（目的地时区）。日期格式为 ISO 8601 注： 预约送货上门仅适用于日本 (JP)。
	PriceDesignation              string      `xml:"PriceDesignation"`               //订单预约送货上门的终止日期（目的地时区）。日期格式为 ISO 8601 注： 预约送货上门仅适用于日本 (JP)。
	IsTransparency                bool        `xml:"IsTransparency"`                 //IsTransparency
	SerialNumberRequired          bool        `xml:"SerialNumberRequired"`           //SerialNumberRequired
	TaxCollectionModel            string      `xml:"TaxCollection>Model"`            //TaxCollection Model
	TaxCollectionResponsibleParty string      `xml:"TaxCollection>ResponsibleParty"` //TaxCollection ResponsibleParty
}

//InvoiceData  发票信息(仅限中国)
type InvoiceData struct {
	InvoiceRequirement           string `xml:"InvoiceRequirement"`           //发票要求信息。 Individual - 买家要求对订单中的每件订单商品单独开具发票。 Consolidated - 买家要求对订单中的所有订单商品开具一张发票。 MustNotSend - 买家不要求开具发票。
	BuyerSelectedInvoiceCategory string `xml:"BuyerSelectedInvoiceCategory"` //买家在下订单时选择的发票类目信息。
	InvoiceTitle                 string `xml:"InvoiceTitle"`                 //买家指定的发票抬头。
	InvoiceInformation           string `xml:"InvoiceInformation"`           //发票信息。 NotApplicable - 买家不- 要求开具发票。 BuyerSelectedInvoiceCategory - 亚马逊建议将此项操作返回的BuyerSelectedInvoiceCategory值作为发票上的发票类目 ProductTitle - 亚马逊建议将商品名称作为发票上的发票类目。
}


// ListOrderItems operation returns order item information for an AmazonOrderId that
// you specify. The order item information includes Title, ASIN, SellerSKU, ItemPrice,
// ShippingPrice, as well as tax and promotion information.  You can retrieve order item
// information by first using the ListOrders operation to find orders created or updated
// during a time frame that you specify. An AmazonOrderId is included with each order that
// is returned. You can then use these AmazonOrderId values with the ListOrderItems
// operation to get detailed order item information for each order.
func (s *OrderService) ListOrderItems(c *Credential, amazonOrderID string) (requestID string, orderItemsResult *OrderItemsResult, err error) {
	data := ActionValues("ListOrderItems")
	data.Set("AmazonOrderId", amazonOrderID)

	var response struct {
		BaseResponse
		OrderItems *OrderItemsResult `xml:"ListOrderItemsResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.OrderItems, nil
}

// ListOrderItemsByNextToken operation returns the next page of order items using
// the NextToken value that was returned by your previous request to either ListOrderItems
// or ListOrderItemsByNextToken. If NextToken is not returned, there are no more pages to return.
func (s *OrderService) ListOrderItemsByNextToken(c *Credential, amazonOrderID, nextToken string) (string, *OrderItemsResult, error) {
	data := ActionValues("ListOrderItemsByNextToken")
	data.Set("NextToken", nextToken)

	var response struct {
		BaseResponse
		OrderItems *OrderItemsResult `xml:"ListOrderItemsByNextTokenResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.OrderItems, nil
}
