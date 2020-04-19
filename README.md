# gositemap

[sitemap 协议](https://www.sitemaps.org/protocol.html)

go语言实现的sitemap生成工具
```go
func main() {
    st := NewSiteMap()
    st.SetPretty(true)

    url := NewUrl()
    url.SetLoc("https://www.douyacun.com/")
    url.SetLastmod(time.Now())
    url.SetChangefreq(Daily)
    url.SetPriority(1)
    st.AppendUrl(url)
    bt, err := st.ToXml()
    if err != nil {
        log.Printf("%v", err)
        return
    }
    log.Printf("%s", bt)
}
```
输出
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://www.douyacun.com/</loc>
    <lastmod>2020-04-19 17:22:41.789997 +0800 CST m=+0.000658489</lastmod>
    <changefreq>daily</changefreq>
    <priority>1</priority>
  </url>
</urlset>
```

# Features
- [ x ] [Image sitemap 图片](#image-sitemap)
- [ x ] [News sitemap 新闻](#news-sitemap)
- [ x ] [Video sitemap 视频](#video-sitemap)
- [ x ] [file storage](#file-storage)

# Image sitemap
```go
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
```

输出
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1">
  <url>
    <loc>https://www.douyacun.com/show_image</loc>
    <image:image>
      <image:loc>https://www.douyacun.com/image1.jpg</image:loc>
      <image:caption>图片的说明。</image:caption>
      <image:title>example</image:title>
      <image:license>https://www.douyacun.com</image:license>
    </image:image>
    <image:image>
      <image:loc>https://www.douyacun.com/image2.jpg</image:loc>
      <image:caption>图片的说明。</image:caption>
      <image:title>example</image:title>
      <image:license>https://www.douyacun.com</image:license>
    </image:image>
  </url>
</urlset>
```

# News sitemap
```go
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
```
输出
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9">
  <url>
    <loc>https://www.douyacun.com/business/article55.html</loc>
    <news:news>
      <news:publication>
        <news:name>《示例时报》</news:name>
        <news:language>zh-cn</news:language>
      </news:publication>
      <news:publication_date>2020-04-19T17:24:49+08:00</news:publication_date>
      <news:title>公司 A 和 B 正在进行合并谈判</news:title>
    </news:news>
  </url>
</urlset>
```

# Video sitemap
```go
st := NewSiteMap()
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
```

输出
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
  <url>
    <loc>https://www.douyacun.com</loc>
    <video:video>
      <video:thumbnail_loc>http://www.example.com/thumbs/123.jpg</video:thumbnail_loc>
      <video:title>适合夏季的烧烤排餐</video:title>
      <video:description>小安教您如何每次都能烤出美味牛排</video:description>
      <video:content_loc>http://streamserver.example.com/video123.mp4</video:content_loc>
      <video:player_loc allow_embed="yes">http://www.example.com/videoplayer.php?video=123</video:player_loc>
      <video:duration>600</video:duration>
      <video:expiration_date>2020-04-20T17:25:49+08:00</video:expiration_date>
      <video:rating>4.2</video:rating>
      <video:view_count>12345</video:view_count>
      <video:publication_date>2020-04-18T17:25:49+08:00</video:publication_date>
      <video:family_friendly>yes</video:family_friendly>
      <video:restriction relationship="allow">IE GB US CN CA</video:restriction>
      <video:price currency="EUR" type="own" resolution="hd">6.99</video:price>
      <video:requires_subscription>yes</video:requires_subscription>
      <video:uploader info="http://www.example.com/users/grillymcgrillerson">GrillyMcGrillerson</video:uploader>
      <video:live>yes</video:live>
    </video:video>
  </url>
</urlset>
```

[google support](https://support.google.com/news/publisher/answer/74288)


# File storage