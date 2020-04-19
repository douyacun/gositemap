package gositemap

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	TooMuchLinksError = errors.New("单个sitemap文件过多")
)

type urlSet struct {
	*base
	XMLName    xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	XMLNSVideo string   `xml:"xmlns:video,attr,omitempty"`
	XMLNSImage string   `xml:"xmlns:image,attr,omitempty"`
	XMLNSNews  string   `xml:"xmlns:news,attr,omitempty"`
	Token      []xml.Token
}

type sitemap struct {
	*urlSet
	*options
}

func NewSiteMap() *sitemap {
	return &sitemap{
		options: NewOptions(),
		urlSet: &urlSet{
			base: &base{},
		},
	}
}

func (s *sitemap) AppendUrl(url *url) {
	if !strings.HasPrefix(url.Loc, "http") {
		url.Loc = strings.TrimRight(s.defaultHost, "/") + strings.TrimLeft(url.Loc, "/")
	}
	s.setNs(url.xmlns)
	s.Token = append(s.Token, url)
}

func (s *sitemap) ToXml() ([]byte, error) {
	if ImageXmlNS&s.xmlns == ImageXmlNS {
		s.urlSet.XMLNSImage = "http://www.google.com/schemas/sitemap-image/1.1"
	}
	if VideoXmlNS&s.xmlns == VideoXmlNS {
		s.urlSet.XMLNSVideo = "http://www.google.com/schemas/sitemap-video/1.1"
	}
	if NewsXmlNS&s.xmlns == NewsXmlNS {
		s.urlSet.XMLNSNews = "http://www.google.com/schemas/sitemap-news/0.9"
	}
	if len(s.urlSet.Token) > s.options.maxLinks {
		return nil, TooMuchLinksError
	}
	var (
		data []byte
		err  error
		buf  bytes.Buffer
	)
	if s.options.pretty {
		buf.Write([]byte(xml.Header))
		data, err = xml.MarshalIndent(s, "", "  ")
	} else {
		buf.Write([]byte(strings.Trim(xml.Header, "\n")))
		data, err = xml.Marshal(s)
	}
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	return buf.Bytes(), nil
}

// filename 生成sitemap文件名
func (s *sitemap) Storage() (filename string, err error) {
	var (
		data []byte
	)
	data, err = s.ToXml()
	if err != nil {
		return
	}
	if err = os.MkdirAll(s.publicPath, 0755); err != nil {
		return
	}
	if s.compress {
		basename := strings.TrimRight(s.filename, path.Ext(s.filename))
		filename = basename + ".xml.gz"
		if fd, err := os.OpenFile(path.Join(s.publicPath, filename), os.O_WRONLY|os.O_CREATE, 0666); err != nil {
			return "", err
		} else {
			gw := gzip.NewWriter(fd)
			if _, err := gw.Write(data); err != nil {
				return "", err
			}
			_ = fd.Close()
			_ = gw.Close()
			return filename, nil
		}
	} else {
		return s.filename, ioutil.WriteFile(path.Join(s.publicPath, s.filename), data, 0666)
	}
}

type Map struct {
	XMLName xml.Name `xml:"sitemap"`
	Loc     string   `xml:"loc"`
}

type siteMapIndex struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 sitemapindex"`
	SiteMap []Map
}

func NewSiteMapIndex() *siteMapIndex {
	return &siteMapIndex{
		SiteMap: make([]Map, 0),
	}
}

func (s *siteMapIndex) Append(loc string) {
	m := Map{
		Loc: loc,
	}
	s.SiteMap = append(s.SiteMap, m)
}

func (s *siteMapIndex) ToXml() ([]byte, error) {
	var (
		data []byte
		err  error
		buf  bytes.Buffer
	)
	buf.Write([]byte(xml.Header))
	data, err = xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	return buf.Bytes(), nil
}

func (s *siteMapIndex) Storage(filepath string) (filename string, err error) {
	if path.Ext(filepath) != ".xml" {
		return "", errors.New("建议以.xml作为文件扩展名")
	}
	var data []byte
	if data, err = s.ToXml(); err != nil {
		return
	}
	err = ioutil.WriteFile(filepath, data, 0666)
	filename = path.Base(filepath)
	return
}
