package gositemap

import (
	"testing"
)

func TestNewSiteMap(t *testing.T) {
	st := NewSiteMap()
	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/")
	url.SetChangefreq(Daily)
	st.AppendUrl(url)
	bt, err := st.ToXml()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%s", bt)
}
