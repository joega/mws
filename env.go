package mws

import (
	"fmt"
	"time"
)

//Env  .
var Env env


func init() {
	//名称
	Env.nameMap = map[string]string{
		"ATVPDKIKX0DER":  "US", // United States
		"A2EUQ1WTGCTBG2": "CA", // Canada
		"A1AM78C64UM0Y8": "MX", // Mexico
		"A2Q3Y263D00KWC": "BR", // Brazil
		"A2VIGQ35RCS4UG": "AE", // United Arab Emirates
		"A1PA6795UKMFR9": "DE", // Germany
		"A1RKKUPIHCS9HS": "ES", // Spain
		"A13V1IB3VIYZZH": "FR", // France
		"A1F83G8C2ARO7P": "UK", // United Kingdom
		"A21TJRUUN4KGV":  "IN", // India
		"APJ6JRA9NG5V4":  "IT", // Italy
		"A33AVAJ2PDY3EV": "TU", // Turkey
		"A39IBJ37TRP1C6": "AU", // Australia
		"A1VC38T7YXB528": "JP", // Japan
		"AAHKV2X7AFYLW":  "CN", // China
		"A19VAU5U5O7RUS": "SG", // Singapore
		"ARBP9OOSHTCHU": "EG", // Egypt
		"A1805IZSGTT6HS": "NL", // Netherlands

	}
	//域名
	Env.webDomainMap = map[string]string{
		"ATVPDKIKX0DER":  "amazon.com",    // US United States
		"A2EUQ1WTGCTBG2": "amazon.ca",     // CA Canada
		"A1AM78C64UM0Y8": "amazon.com.mx", // MX Mexico
		"A2Q3Y263D00KWC": "amazon.com.br", // BR Brazil
		"A2VIGQ35RCS4UG": "amazon.ae",     // AE United Arab Emirates
		"A1PA6795UKMFR9": "amazon.de",     // DE Germany
		"A1RKKUPIHCS9HS": "amazon.es",     // ES Spain
		"A13V1IB3VIYZZH": "amazon.fr",     // FR France
		"A1F83G8C2ARO7P": "amazon.co.uk",  // UK United Kingdom
		"A21TJRUUN4KGV":  "amazon.in",     // IN India
		"APJ6JRA9NG5V4":  "amazon.it",     // IT Italy
		"A33AVAJ2PDY3EV": "amazon.com.tr", // TR Turkey
		"A39IBJ37TRP1C6": "amazon.com.au", // AU Australia
		"A1VC38T7YXB528": "amazon.co.jp",  // JP Japan
		"AAHKV2X7AFYLW":  "amazon.cn",     // CN China
		"A19VAU5U5O7RUS": "amazon.sg", 	   // SG Singapore
		"ARBP9OOSHTCHU": "amazon.eg",      // EG Egypt
		"A1805IZSGTT6HS": "amazon.nl",     // NL Netherlands
	}
	//开发者平台域名
	Env.mwsDomainMap = map[string]string{
		"ATVPDKIKX0DER":  "mws.amazonservices.com",    // US United States
		"A2EUQ1WTGCTBG2": "mws.amazonservices.ca",     // CA Canada
		"A1AM78C64UM0Y8": "mws.amazonservices.com.mx", // MX Mexico
		"A2Q3Y263D00KWC": "mws.amazonservices.com",    // BR Brazil
		"A2VIGQ35RCS4UG": "mws-eu.amazonservices.com", // AE United Arab Emirates
		"A1PA6795UKMFR9": "mws-eu.amazonservices.com", // DE Germany
		"A1RKKUPIHCS9HS": "mws-eu.amazonservices.com", // ES Spain
		"A13V1IB3VIYZZH": "mws-eu.amazonservices.com", // FR France
		"A1F83G8C2ARO7P": "mws-eu.amazonservices.com", // UK United Kingdom
		"APJ6JRA9NG5V4":  "mws-eu.amazonservices.com", // IT Italy
		"A21TJRUUN4KGV":  "mws.amazonservices.in",     // IN India
		"A33AVAJ2PDY3EV": "mws-eu.amazonservices.com", // TR Turkey
		"A1VC38T7YXB528": "mws.amazonservices.jp",     // JP Japan
		"A39IBJ37TRP1C6": "mws.amazonservices.com.au", // AU Australia
		"AAHKV2X7AFYLW":  "mws.amazonservices.com.cn", // CN China
		"A19VAU5U5O7RUS": "mws-fe.amazonservices.com", // SG Singapore
		"ARBP9OOSHTCHU":  "mws-eu.amazonservices.com", // EG Egypt
		"A1805IZSGTT6HS": "mws-eu.amazonservices.com", // NL Netherlands
	}
	//时区 INNA时区
	Env.tzMap = map[string]string{
		"ATVPDKIKX0DER":  "America/Los_Angeles", // US United States
		"A2EUQ1WTGCTBG2": "America/Toronto",     // CA Canada
		"A1AM78C64UM0Y8": "America/Mexico_City", // MX Mexico
		"A2Q3Y263D00KWC": "America/Sao_Paulo",   // BR Brazil
		"A2VIGQ35RCS4UG": "Asia/Dubai",          // AE United Arab Emirates
		"A1PA6795UKMFR9": "Europe/Berlin",       // DE Germany
		"A1RKKUPIHCS9HS": "Europe/Madrid",       // ES Spain
		"A13V1IB3VIYZZH": "Europe/Paris",        // FR France
		"A1F83G8C2ARO7P": "Europe/London",       // UK United Kingdom
		"A21TJRUUN4KGV":  "Asia/Kolkata",        // IN India
		"APJ6JRA9NG5V4":  "Europe/Rome",         // IT Italy
		"A33AVAJ2PDY3EV": "Europe/Istanbul",     // TR Turkey
		"A39IBJ37TRP1C6": "Australia/Sydney",    // AU Australia
		"A1VC38T7YXB528": "Asia/Tokyo",          // JP Japan
		"AAHKV2X7AFYLW":  "Asia/Shanghai",       // CN China
		"A19VAU5U5O7RUS": "Asia/Singapore",      // SG Singapore
		"ARBP9OOSHTCHU":  "Africa/Cairo",        // EG Egypt
		"A1805IZSGTT6HS": "Europe/Amsterdam",    // NL Netherlands
	}
	//国家分组
	Env.areaMap = map[string][]string{
		"NA": {"ATVPDKIKX0DER", "A2EUQ1WTGCTBG2", "A1AM78C64UM0Y8", "A2Q3Y263D00KWC"},
		"EU": {"A2VIGQ35RCS4UG", "A1PA6795UKMFR9", "A1RKKUPIHCS9HS", "A13V1IB3VIYZZH", "A1F83G8C2ARO7P", "A21TJRUUN4KGV", "APJ6JRA9NG5V4", "A33AVAJ2PDY3EV", "ARBP9OOSHTCHU", "A1805IZSGTT6HS" },
		"FE": {"A39IBJ37TRP1C6", "A1VC38T7YXB528", "A19VAU5U5O7RUS"},
		"CN": {"AAHKV2X7AFYLW"},
	}
	//国家关联
	Env.linkMap = map[string][]string{
		"ATVPDKIKX0DER":  Env.areaMap["NA"], // US United States
		"A2EUQ1WTGCTBG2": Env.areaMap["NA"], // CA Canada
		"A1AM78C64UM0Y8": Env.areaMap["NA"], // MX Mexico
		"A2Q3Y263D00KWC": Env.areaMap["NA"], // BR Brazil
		"A2VIGQ35RCS4UG": Env.areaMap["EU"], // AE United Arab Emirates
		"A1PA6795UKMFR9": Env.areaMap["EU"], // DE Germany
		"A1RKKUPIHCS9HS": Env.areaMap["EU"], // ES Spain
		"A13V1IB3VIYZZH": Env.areaMap["EU"], // FR France
		"A1F83G8C2ARO7P": Env.areaMap["EU"], // UK United Kingdom
		"A21TJRUUN4KGV":  Env.areaMap["EU"], // IN India
		"APJ6JRA9NG5V4":  Env.areaMap["EU"], // IT Italy
		"A33AVAJ2PDY3EV": Env.areaMap["EU"], // TR Turkey
		"A39IBJ37TRP1C6": Env.areaMap["FE"], // AU Australia
		"A1VC38T7YXB528": Env.areaMap["FE"], // JP Japan
		"AAHKV2X7AFYLW":  Env.areaMap["CN"], // CN China
		"A19VAU5U5O7RUS": Env.areaMap["FE"], // SG Singapore
		"ARBP9OOSHTCHU":  Env.areaMap["EU"], // EG Egypt
		"A1805IZSGTT6HS": Env.areaMap["EU"], // NL Netherlands
	}
}

type env struct {
	nameMap      map[string]string
	webDomainMap map[string]string
	mwsDomainMap map[string]string
	tzMap        map[string]string
	linkMap      map[string][]string
	areaMap      map[string][]string
}

//GetTimezone 获取站点时区
func (e env) GetTimezone(marketplaceID string) *time.Location {
	innaName := e.tzMap[marketplaceID]
	if innaName != "" {
		l, err := time.LoadLocation(innaName)
		if err == nil {
			return l
		}
	}
	return time.UTC
}

//GetWebDomain 获取站点网站域名
func (e env) GetWebDomain(marketplaceID string) string {
	d := e.webDomainMap[marketplaceID]
	if d == "" {
		d = "amazon.com"
	}
	return d
}

//GetWebURL 获取站点域名前缀(含scheme)
func (e env) GetWebURL(marketplaceID string) string {
	return "https://www." + e.GetWebDomain(marketplaceID)
}

//GetSellercentralURL 获取站点登录域名(含scheme)
func (e env) GetSellercentralURL(marketplaceID string) string {
	return "https://sellercentral." + e.GetWebDomain(marketplaceID)
}

//GetMWSDomain 获取mws域名
func (e env) GetMWSDomain(marketplaceID string) string {
	d := e.mwsDomainMap[marketplaceID]
	if d == "" {
		d = "mws.amazonservices.com"
	}
	return d
}

//GetCountryName 获取站点国家名称
func (e env) GetCountryName(marketplaceID string) string {
	name := e.nameMap[marketplaceID]
	if name == "" {
		name = fmt.Sprintf("未知[%s]", marketplaceID)
	}
	return name
}

//GetLinkedMarketplace 获取关联的商城国家ID
func (e env) GetLinkedMarketplace(marketplaceID string) []string {
	marketplaces := e.linkMap[marketplaceID]
	if len(marketplaces) == 0 {
		marketplaces = []string{marketplaceID}
	}
	return marketplaces
}
