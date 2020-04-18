package gositemap

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

var (
	InvalidLocError         = errors.New("不得与 <loc> 网址相同")
	InvalidDescriptionError = errors.New("视频的说明。不得超过 2048 个字符")
	InvalidDurationError    = errors.New("视频的时长。值必须介于 1（含）和 28800（8 小时，含）之间")
	InvalidRatingError      = errors.New("视频的评分。支持的值为介于 0.0（下限，含）到 5.0（上限，含）之间的浮点数")
	InvalidRestrictionError = errors.New("国家/地区代码不支持")
	InvalidCurrencyError    = errors.New("无效币种")
	InvalidTagError         = errors.New("最多允许使用 32 个标签")
)

type platform string

const (
	Web    platform = "web"
	Mobile platform = "mobile"
	TV     platform = "tv"
)

// 指向特定视频的播放器的网址
// allow_embed [可选] Google 是否可以将视频嵌入搜索结果中。允许的值为 yes 或 no。
type PlayerLoc struct {
	XMLName    xml.Name `xml:"video:player_loc"`
	AllowEmbed string   `xml:"allow_embed,attr"`
	Content    string   `xml:",chardata"`
}

// 在哪些国家展示
// Relationship
// 	- allow: 允许展示
//  - deny: 禁止展示
// 各国对应的code：https://en.wikipedia.org/wiki/ISO_3166-3#Current_codes
type Restriction struct {
	XMLName      xml.Name `xml:"video:restriction"`
	Relationship string   `xml:"relationship,attr"`
	Content      string   `xml:",chardata"`
}

// 是否在指定平台展示，web/mobile/tv
// Relationship
// 	- allow: 允许展示
//  - deny: 禁止展示
type Platform struct {
	XMLName      xml.Name `xml:"video:platform"`
	Relationship string   `xml:"relationship,attr"`
	Content      platform `xml:",chardata"`
}

// 采购价格
type Price struct {
	XMLName    xml.Name `xml:"video:price"`
	Currency   string   `xml:"currency,attr"`   // 货币，https://en.wikipedia.org/wiki/ISO_4217
	Type       string   `xml:"type,attr"`       // 采购方式, rent: 出租，own：拥有
	Resolution string   `xml:"resolution,attr"` // 清晰度, hd,sd
	Content    float64  `xml:",chardata"`
}

// 视频上传者信息
type Uploader struct {
	XMLName xml.Name `xml:"video:price"`
	Info    string   `xml:"info,attr"` // 必须要和视频域名(loc)相同
	Content string   `xml:",chardata"`
}

type Video struct {
	XMLName              xml.Name `xml:"video:video"`
	ThumbnailLoc         string   `xml:"video:thumbnail_loc"` // 视频缩略图文件的网址
	Title                string   `xml:"video:title"`         // 视频标题
	Description          string   `xml:"video:description"`   // 视频的说明，不得超过 2048 个字符
	ContentLoc           string   `xml:"video:content_loc"`   // 指向实际视频媒体文件的网址
	PlayerLoc            *PlayerLoc
	Duration             time.Duration `xml:"video:duration,omitempty"`         // 视频的时长
	ExpirationDate       string        `xml:"video:expiration_date,omitempty"`  // 视频的失效日期
	Rating               int           `xml:"video:rating,omitempty"`           // 视频评分
	ViewCount            int           `xml:"video:view_count,omitempty"`       // 视频观看次数
	PublicationDate      string        `xml:"video:publication_date,omitempty"` // 第一次发布视频的日期
	FamilyFriendly       string        `xml:"video:family_friendly,omitempty"`  // yes/no 开启安全搜索或关闭的情况下播放。
	Restriction          *Restriction
	Platform             *Platform
	Price                *Price
	RequiresSubscription string `xml:"video:requires_subscription,omitempty"` // 是否需要订阅（付费或免费）才能观看视频 yes/no
	Uploader             *Uploader
	Live                 string `xml:"video:live,omitempty"`
	Tag                  string `xml:"video:tag,omitempty"`      // 视频标签，建议短标签，每个tag一个标签，最多支持32个标签
	Category             string `xml:"video:category,omitempty"` // 分类
}

func NewVideo() *Video {
	return &Video{}
}

// 指向视频缩略图文件的网址, 從 160x90 到 1920x1080 像素 建议png、jpg
func (v *Video) SetThumbnailLoc(loc string) {
	v.ThumbnailLoc = loc
}

// 视频标题
func (v *Video) SetTitle(title string) {
	v.Title = title
}

// 视频的说明。不得超过 2048 个字符
func (v *Video) SetDescription(description string) {
	if len(description) > 2048 {
		panic(InvalidDescriptionError)
	}
	v.Description = description
}

// 指向实际视频媒体文件的网址
func (v *Video) SetContentLoc(loc string) {
	if v.PlayerLoc != nil {
		panic(InvalidLocError)
	}
	v.ContentLoc = loc
}

// 指向特定视频的播放器的网址
// embed, 是否可以将视频嵌入搜索结果中
func (v *Video) SetPlayerLoc(loc string, embed bool) {
	if v.ContentLoc != "" {
		panic(InvalidLocError)
	}
	v.PlayerLoc = &PlayerLoc{
		Content:    loc,
		AllowEmbed: "yes",
	}
	if !embed {
		v.PlayerLoc.AllowEmbed = "no"
	}
}

// 视频的时长（以秒为单位）。值必须介于 1（含）和 28800（8 小时，含）之间
func (v *Video) SetDuration(duration time.Duration) {
	if duration > 28800*time.Second {
		panic(InvalidDurationError)
	}
	v.Duration = duration
}

// 视频的失效日期
func (v *Video) SetExpirationDate(date time.Time) {
	v.ExpirationDate = date.Format(time.RFC3339)
}

// 视频的评分。支持的值为介于 0.0（下限，含）到 5.0（上限，含）之间的浮点数
func (v *Video) SetRating(rating int) {
	if rating < 0 || rating > 5 {
		panic(InvalidRatingError)
	}
	v.Rating = rating
}

// 视频的观看次数。
func (v *Video) SetViewCount(count int) {
	v.ViewCount = count
}

// 第一次发布视频的日期
func (v *Video) SetPublicationDate(date time.Time) {
	v.PublicationDate = date.Format(time.RFC3339)
}

// 视频是否在安全搜索的情况下播放。
func (v *Video) SetFamilyFriendly(yes bool) {
	if yes {
		v.FamilyFriendly = "yes"
	} else {
		v.FamilyFriendly = "no"
	}
}

// 是否在来自特定国家/地区的搜索结果中显示或隐藏您的视频
func (v *Video) SetRestriction(code []string, allow bool) {
	all := []string{"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW"}
	access := 0
	for _, i := range all {
		for _, j := range code {
			if j == i {
				access++
			}
		}
	}
	if access != len(code) {
		panic(InvalidRestrictionError)
	}
	v.Restriction = &Restriction{
		Relationship: "",
		Content:      strings.Join(code, " "),
	}
	if allow {
		v.Restriction.Relationship = "allow"
	} else {
		v.Restriction.Relationship = "deny"
	}
	return
}

// 是否在指定类型的平台上的搜索结果中显示或隐藏您的视频
func (v *Video) SetPlatForm(p platform, allow bool) {
	v.Platform = &Platform{
		Relationship: "",
		Content:      p,
	}
	if allow {
		v.Platform.Relationship = "allow"
	} else {
		v.Platform.Relationship = "deny"
	}
}

// currency 货币，https://en.wikipedia.org/wiki/ISO_4217
// own: 采购方式, true 拥有 false 租用
// hd: 清晰度, true 高清 false 标清
func (v *Video) SetPrice(price float64, currency string, own bool, hd bool) {
	all := []string{"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF", "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UYW", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XTS", "XUA", "XXX", "YER", "ZAR", "ZMW", "ZWL", "CNH", "GGP", "IMP", "JEP", "KID", "NIS", "NTD", "PRB", "SLS", "RMB", "TVD", "ZWB", "DASH", "ETH", "VTC", "BCH", "BTC", "XBT", "XLM", "XMR", "XRP", "ZEC", "LTC", "ADF", "ADP", "AFA", "AOK", "AON", "AOR", "ARL", "ARP", "ARA", "ATS", "AZM", "BAD", "BEF", "BGL", "BOP", "BRB", "BRC", "BRN", "BRE", "BRR", "BYB", "BYR", "CSD", "CSK", "CYP", "DDM", "DEM", "ECS", "ECV", "EEK", "ESA", "ESB", "ESP", "FIM", "FRF", "GNE", "GHC", "GQE", "GRD", "GWP", "HRD", "IEP", "ILP", "ILR", "ISJ", "ITL", "LAJ", "LTL", "LUF", "LVL", "MAF", "MCF", "MGF", "MKN", "MLF", "MVQ", "MRO", "MXP", "MZM", "MTL", "NIC", "NLG", "PEH", "PEI", "PLZ", "PTE", "ROL", "RUR", "SDD", "SDP", "SIT", "SKK", "SML", "SRG", "STD", "SUR", "TJR", "TMM", "TPE", "TRL", "UAK", "UGS", "USS", "UYP", "UYN", "VAL", "VEB", "VEF", "XEU", "XFO", "XFU", "YDD", "YUD", "YUN", "YUR", "YUO", "YUG", "YUM", "ZAL", "ZMK", "ZRZ", "ZRN", "ZWC", "ZWD", "ZWN", "ZWR", "ZWL",}
	access := false
	for _, i := range all {
		if i == currency {
			access = true
			break
		}
	}
	if !access {
		panic(InvalidCurrencyError)
	}
	v.Price = &Price{
		Currency:   currency,
		Type:       "own",
		Resolution: "hd",
		Content:    price,
	}
	if !own {
		v.Price.Resolution = "rent"
	}
	if !hd {
		v.Price.Resolution = "sd"
	}
}

// 指明是否需要订阅（收费或免费）才能观看视频
func (v *Video) SetRequiresSubscription(subscription bool) {
	if subscription {
		v.RequiresSubscription = "yes"
	} else {
		v.RequiresSubscription = "no"
	}
}

// uploader, 视频上传者的名称
// info, 包含有关此上传者的其他信息的网页对应的网址, 该网址必须与 <loc> 标记位于同一个网域中
func (v *Video) SetUploader(uploader, info string) {
	v.Uploader = &Uploader{
		Info:    info,
		Content: uploader,
	}
}

// 指明视频是否为直播视频。支持的值为 yes 或 no。
func (v *Video) SetLive(yes bool) {
	if yes {
		v.Live = "yes"
	} else {
		v.Live = "no"
	}
}

// 用于描述视频的任意字符串标记, 最多允许使用 32 个
func (v *Video) SetTag(tags []string) {
	if len(tags) > 32 {
		panic(InvalidTagError)
	}
	v.Tag = strings.Join(tags, " ")
}

// 视频所属宽泛类别的简短说明
func (v *Video) SetCategory(category string) {
	v.Category = category
}
