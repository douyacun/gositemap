package gositemap

import (
	"encoding/xml"
	"strings"
)

type urlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Token   []xml.Token
}

type siteMap struct {
	*urlSet
	op *options
}

func NewSiteMap() *siteMap {
	return &siteMap{
		op: NewOptions(),
		urlSet: &urlSet{},
	}
}

func (s *siteMap) AppendUrl(url *url) {
	if !strings.HasPrefix(url.Loc, "http") {
		url.Loc = strings.TrimRight(s.op.defaultHost, "/") + strings.TrimLeft(url.Loc, "/")
	}
	s.Token = append(s.Token, url)
}

func (s *siteMap) ToXml() ([]byte, error) {
	return xml.MarshalIndent(s, "", " ")
}
