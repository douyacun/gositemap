package gositemap

import "encoding/xml"

type image struct {
	XMLName     xml.Name `xml:"image:image"`
	Loc         string   `xml:"image:loc"`
	Caption     string   `xml:"image:caption,omitempty"`
	GeoLocation string   `xml:"image:geo_location,omitempty"`
	Title       string   `xml:"image:title,omitempty"`
	License     string   `xml:"image:license,omitempty"`
}

func NewImage() *image {
	return &image{}
}

// 图片的网址。某些情况下，图片网址可能与您的主网站不在同一个网域中
// loc是网页的地址，这里是图片的访问路径
func (i *image) SetLoc(loc string) *image {
	i.Loc = loc
	return i
}

// 图片的说明
func (i *image) SetCaption(caption string) *image {
	i.Caption = caption
	return i
}

// 图片的地理位置。例如 <image:geo_location>Limerick, Ireland</image:geo_location>。
func (i *image) SetGeoLocation(geoLocation string) *image {
	i.GeoLocation = geoLocation
	return i
}

// 图片的标题。
func (i *image) SetTitle(title string) *image {
	i.Title = title
	return i
}

// 图片的授权许可所在的网址。
func (i *image) SetLicense(license string) *image {
	i.License = license
	return i
}
