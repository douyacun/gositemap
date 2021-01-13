package gositemap

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSiteMap_ToXml(t *testing.T) {
	st := NewSiteMap()
	st.SetPretty(true)
	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/")
	url.SetLastmod(time.Now())
	url.SetChangefreq(Daily)
	url.SetPriority(float64(3.234) / float64(10))
	st.AppendUrl(url)
	bt, err := st.ToXml()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", bt)
}

func TestNewVideo(t *testing.T) {
	st := NewSiteMap()
	st.SetPretty(true)
	url := NewUrl()
	url.Loc = "https://www.douyacun.com"

	video := NewVideo().
		SetThumbnailLoc("http://www.example.com/thumbs/123.jpg").
		SetTitle("适合夏季的烧烤排餐").
		SetDescription("小安教您如何每次都能烤出美味牛排").
		SetContentLoc("http://streamserver.example.com/video123.mp4").
		SetPlayerLoc("http://www.example.com/videoplayer.php?video=123", true).
		SetDuration(600*time.Second).
		SetExpirationDate(time.Now().Add(time.Hour*24)).
		SetRating(4.2).
		SetViewCount(12345).
		SetPublicationDate(time.Now().Add(-time.Hour*24)).
		SetFamilyFriendly(true).
		SetRestriction([]string{"IE", "GB", "US", "CN", "CA"}, true).
		SetPrice(6.99, "EUR", true, true).
		SetRequiresSubscription(true).
		SetUploader("GrillyMcGrillerson", "http://www.example.com/users/grillymcgrillerson").
		SetLive(true)

	url.AppendVideo(video)
	st.AppendUrl(url)
	bt, err := st.ToXml()
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("%s", bt)
}

func TestNewImage(t *testing.T) {
	st := NewSiteMap()
	st.SetPretty(true)
	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/show_image")
	// 注意这里url.SetLoc设置网页的网址
	// image.SetLoc 设置的是图片访问路径，image的域名和网址域名不一致也可以
	_image := NewImage().
		SetLoc("https://www.douyacun.com/image1.jpg").
		SetTitle("example").
		SetCaption("图片的说明。").
		SetLicense("https://www.douyacun.com")
	url.AppendImage(_image)

	_image = NewImage().
		SetLoc("https://www.douyacun.com/image2.jpg").
		SetTitle("example").
		SetCaption("图片的说明。").
		SetLicense("https://www.douyacun.com")
	url.AppendImage(_image)

	st.AppendUrl(url)
	data, err := st.ToXml()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", data)
}

func TestNewNews(t *testing.T) {
	st := NewSiteMap()
	st.SetPretty(true)

	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/business/article55.html")
	url.AppendNews(NewNews().SetName("《示例时报》").
		SetTitle("公司 A 和 B 正在进行合并谈判").
		SetLanguage("zh-cn").
		SetPublicationDate(time.Now()))
	st.AppendUrl(url)
	data, err := st.ToXml()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", data)
}

func TestSitemap_Storage(t *testing.T) {
	st := NewSiteMap()
	st.SetDefaultHost("https://www.douyacun.com")
	st.SetPretty(true)
	//st.SetMaxLinks(2)
	st.SetPublicPath("/Users/liuning/Documents/github/gositemap")

	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/business/article55.html")
	url.AppendNews(NewNews().SetName("《示例时报》").
		SetTitle("公司 A 和 B 正在进行合并谈判").
		SetLanguage("zh-cn").
		SetPublicationDate(time.Now()))
	st.AppendUrl(url)

	//url = NewUrl().SetLoc("https://www.douyacun.com/business/article56.html")
	//
	//url.AppendNews(NewNews().SetName("《示例时报》").
	//	SetTitle("公司 A 和 C 正在进行合并谈判").
	//	SetLanguage("zh-cn").
	//	SetPublicationDate(time.Now()))
	//st.AppendUrl(url)
	//
	//url = NewUrl().SetLoc("https://www.douyacun.com/business/article57.html")
	//
	//url.AppendNews(NewNews().SetName("《示例时报》").
	//	SetTitle("公司 A 和 C 正在进行合并谈判").
	//	SetLanguage("zh-cn").
	//	SetPublicationDate(time.Now()))
	//st.AppendUrl(url)

	path, err := st.Storage()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(path)
}

func TestNewSiteMapIndex(t *testing.T) {
	mapIndex := NewSiteMapIndex()

	st1 := NewSiteMap()
	st1.SetPretty(true)
	st1.SetFilename("sitemap1")
	st1.SetCompress(true)

	url := NewUrl()
	url.SetLoc("https://www.douyacun.com/business/article55.html")
	url.AppendNews(NewNews().SetName("《示例时报》").
		SetTitle("公司 A 和 B 正在进行合并谈判").
		SetLanguage("zh-cn").
		SetPublicationDate(time.Now()))
	st1.AppendUrl(url)

	url = NewUrl().SetLoc("https://www.douyacun.com/business/article56.html")

	url.AppendNews(NewNews().SetName("《示例时报》").
		SetTitle("公司 A 和 C 正在进行合并谈判").
		SetLanguage("zh-cn").
		SetPublicationDate(time.Now()))
	st1.AppendUrl(url)
	st1Filename, err := st1.Storage()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	mapIndex.Append("https://www.douyacun.com/" + st1Filename)
	filename, err := mapIndex.Storage("/Users/liuning/Documents/github/gositemap/sitemap_index.xml")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v", filename)
}
