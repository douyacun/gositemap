package gositemap

import "encoding/xml"



type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	xml.Token
}

