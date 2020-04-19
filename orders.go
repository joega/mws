package mws

import "time"


//OrderService
type OrderService struct {
	*Client
}

//OrdersResult 订单结果
type OrdersResult struct {
	NextToken         string
	LastUpdatedBefore time.Time
	CreatedBefore     time.Time
	Orders            []*Order `xml:"Orders>Order"`
}

//Order  订单模型
type Order struct {
	AmazonOrderID                string                        `xml:"AmazonOrderId"`                //亚马逊所定义的订单编码，格式为 3-7-7。
	SellerOrderID                string                        `xml:"SellerOrderId"`                //卖家所定义的订单编码。
	PurchaseDate                 *time.Time                    `xml:"PurchaseDate"`                 //创建订单的日期。
	LastUpdateDate               *time.Time                    `xml:"LastUpdateDate"`               //订单的最后更新日期
	OrderStatus                  string                        `xml:"OrderStatus"`                  //当前的订单状态。
	FulfillmentChannel           string                        `xml:"FulfillmentChannel"`           //订单配送方式：亚马逊配送 (AFN) 或卖家自行配送 (MFN)。
	SalesChannel                 string                        `xml:"SalesChannel"`                 //订单中第一件商品的销售渠道。
	OrderChannel                 string                        `xml:"OrderChannel"`                 //订单中第一件商品的订单渠道。
	ShipServiceLevel             string                        `xml:"ShipServiceLevel"`             //货件服务水平。
	ShippingAddress              Address                       `xml:"ShippingAddress"`              //订单的配送地址。
	OrderTotal                   Money                         `xml:"OrderTotal"`                   //订单的总费用。
	NumberOfItemsShipped         int                           `xml:"NumberOfItemsShipped"`         //已配送的商品数量。
	NumberOfItemsUnshipped       int                           `xml:"NumberOfItemsUnshipped"`       //未配送的商品数量。
	PaymentMethod                string                        `xml:"PaymentMethod"`                //订单的主要付款方式。COD - 货到付款。仅适用于中国 (CN) 和日本 (JP)。CVS - 便利店。仅适用于日本 (JP)。Other - COD 和 CVS 之外的付款方式。注： 可使用多种次级付款方式为 PaymentMethod = COD的订单付款。每种次级付款方式均表示为 PaymentExecutionDetailItem 对象。
	MarketplaceID                string                        `xml:"MarketplaceId"`                //订单生成所在商城的匿名编码。
	BuyerEmail                   string                        `xml:"BuyerEmail"`                   //买家的匿名电子邮件地址。
	BuyerName                    string                        `xml:"BuyerName"`                    //买家姓名。
	BuyerCounty                  string                        `xml:"BuyerCounty"`                  //This element is used only in the Brazil marketplace.
	ShipmentServiceLevelCategory string                        `xml:"ShipmentServiceLevelCategory"` //订单的配送服务级别分类。 Expedited, FreeEconomy, NextDay, SameDay, SecondDay, Scheduled, Standard
	ShippedByAmazonTFM           bool                          `xml:"ShippedByAmazonTFM"`           //指明订单配送方是否是亚马逊配送 (Amazon TFM) 服务。
	TFMShipmentStatus            string                        `xml:"TFMShipmentStatus"`            //亚马逊 TFM订单的状态。仅当ShippedByAmazonTFM = True时返回。请注意：即使当 ShippedByAmazonTFM = True 时，如果您还没有创建货件，也不会返回 TFMShipmentStatus。值： PendingPickUp, LabelCanceled, PickedUp, AtDestinationFC, Delivered, RejectedByBuyer, Undeliverable, ReturnedToSeller 注： 亚马逊 TFM 仅适用于中国 (CN)。
	CbaDisplayableShippingLabel  string                        `xml:"CbaDisplayableShippingLabel"`  //卖家自定义的配送方式，属于Checkout by Amazon (CBA) 所支持的四种标准配送设置中的一种。有关更多信息，请参阅您所在商城 Amazon Payments 帮助中心的“设置灵活配送方式”主题。 注： CBA 仅适用于美国 (US)、英国 (UK) 和德国 (DE) 的卖家。
	OrderType                    string                        `xml:"OrderType"`                    //订单类型。StandardOrder - 包含当前有库存商品的订单。Preorder -所含预售商品（发布日期晚于当前日期）的订单。 注： Preorder 仅在日本 (JP) 是可行的OrderType 值。
	EarliestShipDate             *time.Time                    `xml:"EarliestShipDate"`             //您承诺的订单发货时间范围的第一天。日期格式为 ISO 8601。 仅对卖家配送网络 (MFN) 订单返回。
	LatestShipDate               *time.Time                    `xml:"LatestShipDate"`               //您承诺的订单发货时间范围的最后一天。日期格式为 ISO 8601。对卖家配送网络 (MFN)	和亚马逊物流 (AFN) 订单返回。
	EarliestDeliveryDate         *time.Time                    `xml:"EarliestDeliveryDate"`         //您承诺的订单送达时间范围的第一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	LatestDeliveryDate           *time.Time                    `xml:"LatestDeliveryDate"`           //您承诺的订单送达时间范围的最后一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	IsBusinessOrder              bool                          `xml:"IsBusinessOrder"`              //true if the order is an Amazon Business order. An Amazon Business order is an order where the buyer is a Verified Business Buyer and the seller is an Amazon Business Seller. For more information about the Amazon Business Seller Program
	PurchaseOrderNumber          string                        `xml:"PurchaseOrderNumber"`          //	he purchase order (PO) number entered by the buyer at checkout.
	IsPrime                      bool                          `xml:"IsPrime"`                      //true if the order is a seller-fulfilled Amazon Prime order.
	IsPremiumOrder               bool                          `xml:"IsPremiumOrder"`               //true if the order has a Premium Shipping Service Level Agreement. For more information about Premium Shipping orders, see "Premium Shipping Options" in the Seller Central Help for your marketplace.
	ReplaceOrderID               string                        `xml:"ReplaceOrderId"`               //The AmazonOrderId value for the order that is being replaced.
	IsReplacementOrder           bool                          `xml:"IsReplacementOrder"`           //true if this is a replacement order.
	PaymentExecutionDetail       []*PaymentExecutionDetailItem `xml:"PaymentExecutionDetail"`
	/*	货到付款 (COD) 订单的次级付款方式的相关信息。
		COD 订单是带有PaymentMethod = COD 的订单。
		包含一个或多个PaymentExecutionDetailItem响应元素。
		注： 对于使用某一次级付款方式付款的 COD 订单，将返回一个 PaymentExecutionDetailItem响应元素，
		其中 PaymentExecutionDetailItem/PaymentMethod = COD。
		对于使用多种次级付款方式付款的 COD 订单，将返回两个或多个 PaymentExecutionDetailItem响应元素。
		可选。仅对 COD 订单返回。仅适用于中国 (CN) 和日本 (JP)。*/
}

//Address  地址信息
type Address struct {
	Name          string `xml:"Name"`          //名称。
	AddressLine1  string `xml:"AddressLine1"`  //街道地址。
	AddressLine2  string `xml:"AddressLine2"`  //街道地址。
	AddressLine3  string `xml:"AddressLine3"`  //街道地址。
	City          string `xml:"City"`          //城市。
	County        string `xml:"County"`        //区县。
	District      string `xml:"District"`      //区。
	StateOrRegion string `xml:"StateOrRegion"` //省/自治区/直辖市或地区。
	PostalCode    string `xml:"PostalCode"`    //邮政编码。
	CountryCode   string `xml:"CountryCode"`   //两位数国家/地区代码。格式为 ISO 3166-1-alpha 2 。
	Phone         string `xml:"Phone"`         //电话号码。
	AddressType   string `xml:"AddressType"`   //AddressType
}

//PaymentExecutionDetailItem  付款信息
type PaymentExecutionDetailItem struct {
	Payment       Money  `xml:"Payment"`
	PaymentMethod string `xml:"PaymentMethod"`
}

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




//Orders
func Orders() *OrderService {
	return &OrderService{
		Client: newClient("/Orders/2013-09-01", "2013-09-01"),
	}
}



// ListOrders operation returns a list of orders created or updated during a time
// frame that you specify. You define that time frame using the CreatedAfter parameter
// or the LastUpdatedAfter parameter. You must use one of these parameters, but not
// both. You can also apply a range of filtering criteria to narrow the list of orders
// that is returned. The ListOrders operation includes order information for each order
// returned, including AmazonOrderId, OrderStatus, FulfillmentChannel, and LastUpdateDate.
func (s *OrderService) ListOrders(c *Credential, marketplaces []string, startTime time.Time, timeIsUpdate bool, maxPerPage int64, params ...Values) (string, *OrdersResult, error) {
	data := ActionValues("ListOrders")
	if !startTime.IsZero() {
		if timeIsUpdate {
			data.SetTime("LastUpdatedAfter", startTime)
		} else {
			data.SetTime("CreatedAfter", startTime)
		}
	}
	if maxPerPage > 0 && maxPerPage != 100 {
		data.SetInt("MaxResultsPerPage", maxPerPage)
	}
	data.Sets("MarketplaceId.Id", marketplaces...)
	data.SetAll(params...)

	var response struct {
		BaseResponse
		Orders *OrdersResult `xml:"ListOrdersResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.Orders, nil
}

// ListOrdersByNextToken operation returns the next page of orders using the NextToken
// value that was returned by your previous request to either ListOrders or ListOrdersByNextToken.
// If NextToken is not returned, there are no more pages to return.
func (s *OrderService) ListOrdersByNextToken(c *Credential, nextToken string) (string, *OrdersResult, error) {
	data := ActionValues("ListOrdersByNextToken")
	data.Set("NextToken", nextToken)

	var response struct {
		BaseResponse
		Orders *OrdersResult `xml:"ListOrdersByNextTokenResult"`
	}
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.Orders, nil
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





