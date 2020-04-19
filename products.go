package mws

//ProductService 商品服务
type ProductService struct {
	*Client
}

//GetMyPriceForASINResponse 商品价格获取响应体
type GetMyPriceForASINResponse struct {
	BaseResponse
	GetMyPriceForASINResult []GetMyPriceForASINResult
}

//GetMyPriceForASINResult GetMyPriceForASINResult
type GetMyPriceForASINResult struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN标识
	OfferPrice      []ProductOffer  `xml:"Product>Offers>Offer"`        //供应价格（销售价格）
}

//GetMyPriceForSKUResponse 商品价格获取响应体
type GetMyPriceForSKUResponse struct {
	BaseResponse
	GetMyPriceForSKUResult []GetMyPriceForSKUResult
}

//GetMyPriceForSKUResult GetMyPriceForASINResult
type GetMyPriceForSKUResult struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN标识
	SKUIdentifier   SKUIdentifier   `xml:"Identifiers>SKUIdentifier"`   //商品SKU标识
	OfferPrice      []ProductOffer  `xml:"Product>Offers>Offer"`        //供应价格（销售价格）
}

//ProductOffer ProductPrice
type ProductOffer struct {
	BuyingPrice        ProductBuyingPrice //包含价格信息（包括进行促销的商品）以及运费。
	RegularPrice       Money              //商品的当前价格（不包括进行促销的商品）。不包括运费。
	FulfillmentChannel string             //商品的配送渠道。Amazon - 亚马逊物流。Merchant - 卖家自行配送。
	ItemCondition      string             //商品的状况。有效值：New、Used、Collectible、Refurbished、Club
	ItemSubCondition   string             //商品的子状况(成色)。有效值：New、Mint、Very Good、Good、Acceptable、Poor、Club、OEM、Warranty、Refurbished Warranty、Refurbished、Open Box 或 Other。
	SellerSKU          string             //商品的 SellerSKU。
	SellerID           string             `xml:"SellerId"` //在操作中提交的 SellerId。
}

//ProductBuyingPrice BuyingPrice
type ProductBuyingPrice struct {
	ListingPrice Money //商品的当前价格（包括进行促销的商品）。
	LandedPrice  Money //ListingPrice + Shipping - Points.请注意，如果未返回到岸价格，则上市价格代表具有最低到岸价格的产品。
	Shipping     Money //商品的运费。
	Points       Money //购买商品时提供的亚马逊积分数量及其货币价值。请注意，Points元素仅在日本（JP）返回。
}

//SKUIdentifier SKUIdentifier
type SKUIdentifier struct {
	MarketplaceID string `xml:"MarketplaceId"`
	SellerID      string `xml:"SellerId"`
	SellerSKU     string
}

//GetMatchingProductForIDResponse 商品信息获取响应体
type GetMatchingProductForIDResponse struct {
	BaseResponse
	ProductResults []GetMatchingProductForIDResult `xml:"GetMatchingProductForIdResult"`
}

//GetMatchingProductForIDResult GetMatchingProductForIdResult
type GetMatchingProductForIDResult struct {
	BaseResponse
	ID       string     `xml:"Id,attr"`
	IDType   string     `xml:"IdType,attr"`
	Status   string     `xml:"status,attr"`
	Products []*Product `xml:"Products>Product"`
}

//Product 商品信息模型
type Product struct {
	MarketplaceASIN                      MarketplaceASIN  `xml:"Identifiers>MarketplaceASIN"`                                       //商品ASIN 商城
	Binding                              string           `xml:"AttributeSets>ItemAttributes>Binding"`                              //
	Brand                                string           `xml:"AttributeSets>ItemAttributes>Brand"`                                //商品品牌
	Color                                string           `xml:"AttributeSets>ItemAttributes>Color"`                                //型号，颜色
	Department                           string           `xml:"AttributeSets>ItemAttributes>Department"`                           //
	Feature                              string           `xml:"AttributeSets>ItemAttributes>Feature"`                              //功能
	ItemDimensions                       Dimension        `xml:"AttributeSets>ItemAttributes>ItemDimensions"`                       //商品尺寸
	IsAdultProduct                       bool             `xml:"AttributeSets>ItemAttributes>IsAdultProduct"`                       //
	IsAutographed                        bool             `xml:"AttributeSets>ItemAttributes>IsAutographed"`                        //
	IsEligibleForTradeIn                 bool             `xml:"AttributeSets>ItemAttributes>IsEligibleForTradeIn"`                 //
	IsMemorabilia                        bool             `xml:"AttributeSets>ItemAttributes>IsMemorabilia"`                        //
	IssuesPerYear                        bool             `xml:"AttributeSets>ItemAttributes>IssuesPerYear"`                        //
	Label                                string           `xml:"AttributeSets>ItemAttributes>Label"`                                //标签
	ListPrice                            Money            `xml:"AttributeSets>ItemAttributes>ListPrice"`                            //生产厂商
	Manufacturer                         string           `xml:"AttributeSets>ItemAttributes>Manufacturer"`                         //
	ManufacturerPartsWarrantyDescription string           `xml:"AttributeSets>ItemAttributes>ManufacturerPartsWarrantyDescription"` //厂商保修描述
	Warranty                             string           `xml:"AttributeSets>ItemAttributes>Warranty"`                             //保修描述
	MaterialType                         string           `xml:"AttributeSets>ItemAttributes>MaterialType"`                         //
	Model                                string           `xml:"AttributeSets>ItemAttributes>Model"`                                //型号
	PackageDimensions                    Dimension        `xml:"AttributeSets>ItemAttributes>PackageDimensions"`                    //包装信息
	PackageQuantity                      int              `xml:"AttributeSets>ItemAttributes>PackageQuantity"`                      //包装数量
	PartNumber                           string           `xml:"AttributeSets>ItemAttributes>PartNumber"`                           //
	ProductGroup                         string           `xml:"AttributeSets>ItemAttributes>ProductGroup"`                         //
	ProductTypeName                      string           `xml:"AttributeSets>ItemAttributes>ProductTypeName"`                      //
	Publisher                            string           `xml:"AttributeSets>ItemAttributes>Publisher"`                            //发布者
	SmallImage                           Image            `xml:"AttributeSets>ItemAttributes>SmallImage"`                           //商品图片（小图）
	Studio                               string           `xml:"AttributeSets>ItemAttributes>Studio"`                               //商品标题
	Title                                string           `xml:"AttributeSets>ItemAttributes>Title"`                                //商品标题
	VariationParent                      MarketplaceASIN  `xml:"Relationships>VariationParent>Identifiers>MarketplaceASIN"`         //父Asin信息
	VariationChild                       []VariationChild `xml:"Relationships>VariationChild"`                                      //子Asin信息
	SalesRankings                        []SalesRank      `xml:"SalesRankings>SalesRank"`                                           //排行信息
}

//SalesRank 分类排行
type SalesRank struct {
	ProductCategoryID string `xml:"ProductCategoryId"`
	Rank              int    `xml:"Rank"`
}

//VariationChild VariationChild
type VariationChild struct {
	Color           string          `xml:"Color"`
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"`
}

//Dimension 包装信息
type Dimension struct {
	Height UnitValue `xml:"Height"`
	Width  UnitValue `xml:"Width"`
	Length UnitValue `xml:"Length"`
	Weight UnitValue `xml:"Weight"`
}

//UnitValue 带单位的数量
type UnitValue struct {
	Unit  string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

//Image 图片
type Image struct {
	URL    string    `xml:"URL"`
	Width  UnitValue `xml:"Width"`
	Height UnitValue `xml:"Height"`
}

//MarketplaceASIN 商城模型
type MarketplaceASIN struct {
	MarketplaceID string `xml:"MarketplaceId"`
	ASIN          string `xml:"ASIN"`
}







// GetMyPriceForSKU operation returns pricing information for your own active offer
// listings, based on the ASIN mapped to the SellerSKU and MarketplaceId that you specify.
// Note that if you submit a SellerSKU for a product for which you don’t have an active
// offer listing, the operation returns an empty Offers element. This operation returns
// pricing information for a maximum of 20 offer listings.
// Maximum request quota	Restore rate				Hourly request quota
// 20 requests				10 items every second		36000 requests per hour
func (s *ProductService) GetMyPriceForSKU(c *Credential, marketplace string, sellerSKUList []string, itemCondition string) (requestID string, prices []GetMyPriceForSKUResult, err error) {
	data := ActionValues("GetMyPriceForSKU")
	data.Set("MarketplaceId", marketplace)
	data.Sets("SellerSKUList.SKU", sellerSKUList...)
	data.Set("ItemCondition", itemCondition)
	var result GetMyPriceForSKUResponse
	if err = s.GetModel(c, data, &result); err != nil {
		return
	}
	requestID = result.RequestID
	prices = result.GetMyPriceForSKUResult
	return
}

// GetMyPriceForASIN operation is the same as the GetMyPriceForSKU operation except
// that it uses a MarketplaceId and an ASIN to uniquely identify a product, and it does
// not return the SKUIdentifier element.
func (s *ProductService) GetMyPriceForASIN(c *Credential, marketplace string, asinList []string, itemCondition string) (requestID string, prices []GetMyPriceForASINResult, err error) {
	data := ActionValues("GetMyPriceForASIN")
	data.Set("MarketplaceId", marketplace)
	data.Sets("ASINList.ASIN", asinList...)
	data.Set("ItemCondition", itemCondition)
	var result GetMyPriceForASINResponse
	if err = s.GetModel(c, data, &result); err != nil {
		return
	}
	requestID = result.RequestID
	prices = result.GetMyPriceForASINResult
	return
}



//Products 创建商品服务
func Products() *ProductService {
	return &ProductService{
		Client: newClient("/Products/2011-10-01", "2011-10-01"),
	}
}



// GetMatchingProductForId operation returns a list of products and their attributes,
// based on a list of product identifier values that you specify. Possible product
// identifiers are ASIN, GCID, SellerSKU, UPC, EAN, ISBN, and JAN.
func (s *ProductService) GetMatchingProductForID(c *Credential, marketplaceID string, idType string, idList ...string) (string, []*Product, error) {
	data := ActionValues("GetMatchingProductForId")
	data.Set("MarketplaceId", marketplaceID)
	data.Set("IdType", idType)
	data.Sets("IdList.Id", idList...)

	var response GetMatchingProductForIDResponse
	if err := s.GetModel(c, data, &response); err != nil {
		return "", nil, err
	}
	var products []*Product
	for _, result := range response.ProductResults {
		if result.Status == "Success" {
			products = append(products, result.Products...)
		}
	}
	return response.RequestID, products, nil
}

