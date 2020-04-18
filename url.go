package gositemap

import (
	"encoding/xml"
	"time"
)

type changefreq string

type InvalidPriorityError struct {
	msg string
}

func (e *InvalidPriorityError) Error() string {
	return e.msg
}

const (
	Always  changefreq = "always"
	Hourly  changefreq = "hourly"
	Daily   changefreq = "daily"
	Weekly  changefreq = "weekly"
	Monthly changefreq = "monthly"
	Yearly  changefreq = "yearly"
	Never   changefreq = "never"
)

type url struct {
	XMLName    xml.Name   `xml:"url"`
	Loc        string     `xml:"loc"`
	Lastmod    string     `xml:"lastmod,omitempty"`
	Changefreq changefreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
	Token      []xml.Token
}

func NewUrl() *url {
	return &url{
		Loc:        "",
		Lastmod:    "",
		Changefreq: "",
		Priority:   0,
	}
}

// 网址
func (u *url) SetLoc(loc string) {
	u.Loc = loc
}

// 最后一次修改时间
func (u *url) SetLastmod(lastMod time.Time) {
	u.Lastmod = lastMod.String()
}

// 更新频率
func (u *url) SetChangefreq(freq changefreq) {
	u.Changefreq = freq
}

// 网页优先级
func (u *url) SetPriority(priority float32) {
	if priority < 0 || priority > 1 {
		panic(InvalidPriorityError{"Valid values range from 0.0 to 1.0"})
	}
	u.Priority = priority
}

// 对于单个网页上的多个视频，为该网页创建一个 <loc> 标记，并为该网页上的每个视频创建一个子级 <video> 元素。
func (u *url) AppendVideo(video Video) {
	u.Token = append(u.Token, video)
}

// 对于单个网页上的多个图片，每个 <url> 标记最多可包含 1000 个 <image:image> 标记。
func (u *url) AppendImage(image Image) {
	u.Token = append(u.Token, image)
}
