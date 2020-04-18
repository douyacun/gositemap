package gositemap

import (
	"encoding/xml"
	"errors"
	"time"
)

var (
	InvalidLanguageError = errors.New("语言代码错误")
)

type news struct {
	XMLName         xml.Name `xml:"news:news"`
	Name            string   `xml:"news:publication>news:name"`
	Language        string   `xml:"news:publication>news:language"`
	PublicationDate string   `xml:"news:publication_date"`
	Title           string   `xml:"news:title"`
}

// ISO 639 语言代码 http://www.loc.gov/standards/iso639-2/php/code_list.php

func NewNews() *news {
	return &news{}
}

// 新闻出版物的名称
func (n *news) SetName(name string) *news {
	n.Name = name
	return n
}

func (n *news) SetLanguage(language string) *news {
	if language != "zh-cn" && language != "zh-tw" {
		all := []string{"aa", "ab", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "ba", "bm", "eu", "be", "bn", "bh", "bi", "bo", "bs", "br", "bg", "my", "ca", "cs", "ch", "ce", "zh", "cu", "cv", "kw", "co", "cr", "cy", "cs", "da", "de", "dv", "nl", "dz", "el", "en", "eo", "et", "eu", "ee", "fo", "fa", "fj", "fi", "fr", "fr", "fy", "ff", "ka", "de", "gd", "ga", "gl", "gv", "el", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hr", "hu", "hy", "ig", "is", "io", "ii", "iu", "ie", "ia", "id", "ik", "is", "it", "jv", "ja", "kl", "kn", "ks", "ka", "kr", "kk", "km", "ki", "rw", "ky", "kv", "kg", "ko", "kj", "ku", "lo", "la", "lv", "li", "ln", "lt", "lb", "lu", "lg", "mk", "mh", "ml", "mi", "mr", "ms", "mk", "mg", "mt", "mn", "mi", "ms", "my", "na", "nv", "nr", "nd", "ng", "ne", "nl", "nn", "nb", "no", "ny", "oc", "oj", "or", "om", "os", "pa", "fa", "pi", "pl", "pt", "ps", "qu", "rm", "ro", "ro", "rn", "ru", "sg", "sa", "si", "sk", "sk", "sl", "se", "sm", "sn", "sd", "so", "st", "es", "sq", "sc", "sr", "ss", "su", "sw", "sv", "ty", "ta", "tt", "te", "tg", "tl", "th", "bo", "ti", "to", "tn", "ts", "tk", "tr", "tw", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "cy", "wa", "wo", "xh", "yi", "yo", "za", "zh", "zu",}
		for _, v := range all {
			if v == language {
				n.Language = v
				break
			}
		}
		panic(InvalidLanguageError)
	} else {
		n.Language = language
	}
	return n
}

func (n *news) SetPublicationDate(date time.Time) *news {
	n.PublicationDate = date.Format(time.RFC3339)
	return n
}

func (n *news) SetTitle(title string) *news {
	n.Title = title
	return n
}
