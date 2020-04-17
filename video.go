package gositemap

import (
	"encoding/xml"
	"time"
)

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
	Content      string   `xml:",chardata"`
}

// 采购价格
type Price struct {
	XMLName    xml.Name `xml:"video:price"`
	Currency   string   `xml:"currency,attr"`   // 货币，https://en.wikipedia.org/wiki/ISO_4217
	Type       string   `xml:"type,attr"`       // 采购方式, rent: 出租，own：拥有
	Resolution string   `xml:"resolution,attr"` // 清晰度, hd,sd
	Content    string   `xml:",chardata"`
}

// 视频上传者信息
type Uploader struct {
	XMLName xml.Name `xml:"video:price"`
	Info    string   `xml:"info,attr"` // 必须要和视频域名(loc)相同
	Content string   `xml:",chardata"`
}

type Video struct {
	XMLName              xml.Name  `xml:"video:video"`
	ThumbnailLoc         string    `xml:"video:thumbnail_loc"`
	Title                string    `xml:"video:title"`
	Description          string    `xml:"video:description"`
	ContentLoc           string    `xml:"video:content_loc"`
	PlayerLoc            string    `xml:"video:player_loc"`
	Duration             int       `xml:"video:duration,omitempty"`
	ExpirationDate       time.Time `xml:"video:expiration_date,omitempty"`
	Rating               int       `xml:"video:rating,omitempty"`
	ViewCount            int       `xml:"video:view_count,omitempty"`
	PublicationDate      int       `xml:"video:publication_date,omitempty"`
	FamilyFriendly       string    `xml:"video:family_friendly,omitempty"` //
	Restriction          *Restriction
	Price                *Price
	RequiresSubscription string `xml:"video:requires_subscription,omitempty"` // 是否需要订阅（付费或免费）才能观看视频 yes/no
	Uploader             *Uploader
	Tag                  string `xml:"video:tag,omitempty"`      // 视频标签，建议短标签，每个tag一个标签，最多支持32个标签
	Category             string `xml:"video:category,omitempty"` // 分类
}
