package gositemap

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

func (s *sitemap) Storage() (string, error) {
	// 5w个连接生成一个文件
	length := len(s.Token)
	if length < MaxSitemapLinks {
		data, err := s.ToXml()
		if err != nil {
			return "", err
		}
		if err = os.MkdirAll(s.publicPath, 0755); err != nil {
			return "", err
		}
		return s.filename, ioutil.WriteFile(path.Join(s.publicPath, s.filename), data, 0666)
	} else {
		index := 1
		basename := strings.TrimRight(s.filename, path.Ext(s.filename))
		si := NewSiteMapIndex()
		host := strings.TrimRight(s.defaultHost, "/")
		publicPath := strings.TrimRight(s.publicPath, "/")
		for i := 0; i < len(s.Token); i += MaxSitemapLinks {
			st := NewSiteMap()
			st.options = s.options
			st.Token = s.Token[i : i+MaxSitemapLinks]
			data, err := st.ToXml()
			filename := fmt.Sprintf("%s%d.xml.gz", basename, index)
			fileAbsPath := path.Join(s.publicPath, filename)
			fd, err := os.OpenFile(fileAbsPath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				return "", err
			}
			gw := gzip.NewWriter(fd)
			if _, err := gw.Write(data); err != nil {
				return "", err
			}
			_ = gw.Close()
			_ = fd.Close()
			index++
			si.Append(NewMap(host + "/" + filename))
		}
		data, err := si.ToXml()
		if err != nil {
			return "", err
		}
		filename := basename + "_index.xml"
		file := publicPath + "/" + filename
		err = ioutil.WriteFile(file, data, 0666)
		return file, err
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
		SiteMap: make([]Map, 1),
	}
}

func (s *siteMapIndex) Append(m Map) {
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

func NewMap(loc string) Map {
	return Map{
		Loc: loc,
	}
}
