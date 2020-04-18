package gositemap

import "encoding/xml"

type Image struct {
	XMLName     xml.Name `xml:"image:image"`
	Loc         string   `xml:"image:loc"`
	Caption     string   `xml:"image:caption,omitempty"`
	GeoLocation string   `xml:"image:geo_location,omitempty"`
	Title       string   `xml:"image:title,omitempty"`
	License     string   `xml:"image:license,omitempty"`
}

func NewImage() *Image {
	return &Image{}
}

// 图片的网址。某些情况下，图片网址可能与您的主网站不在同一个网域中
// loc是网页的地址，这里是图片的访问路径
func (i *Image) SetLoc(loc string) {
	i.Loc = loc
}

// 图片的说明
func (i *Image) SetCaption(caption string) {
	i.Caption = caption
}

// 图片的地理位置。例如 <image:geo_location>Limerick, Ireland</image:geo_location>。
func (i *Image) SetGeoLocation(geoLocation string) {
	i.GeoLocation = geoLocation
}

// 图片的标题。
func (i *Image) SetTitle(title string) {
	i.Title = title
}

// 图片的授权许可所在的网址。
func (i *Image) SetLicense(license string) {
	i.License = license
}
