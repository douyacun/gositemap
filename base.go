package gositemap

type xmlns int8

const (
	ImageXmlNS xmlns = 1
	VideoXmlNS xmlns = 2
	NewsXmlNS  xmlns = 4
)

type base struct {
	xmlns xmlns
}

// image: 0001
// video: 0010
//  news: 0100
func (b *base) setNs(xmlns xmlns) {
	b.xmlns = b.xmlns | xmlns
}
