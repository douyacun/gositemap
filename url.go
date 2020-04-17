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

type Url struct {
	XMLName    xml.Name   `xml:"url"`
	Loc        string     `xml:"loc"`
	Lastmod    string     `xml:"lastmod,omitempty"`
	Changefreq changefreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
	xml.Token
}

func NewUrl() *Url {
	return &Url{
		Loc:        "",
		Lastmod:    "",
		Changefreq: "",
		Priority:   0,
	}
}

func (u *Url) SetLoc(loc string) {
	u.Loc = loc
}

func (u *Url) SetLastmod(lastMod time.Time) {
	u.Lastmod = lastMod.String()
}

func (u *Url) SetChangefreq(freq changefreq) {
	u.Changefreq = freq
}

func (u *Url) SetPriority(priority float32) {
	if priority < 0 || priority > 1 {
		panic(InvalidPriorityError{"Valid values range from 0.0 to 1.0"})
	}
	u.Priority = priority
}

func (u *Url) SetVideo(video Video) {
	u.Token = []Video{
		video,
	}
}
