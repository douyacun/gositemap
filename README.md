# gositemap

[sitemap 协议](https://www.sitemaps.org/protocol.html)

go语言实现的sitemap生成工具
```go
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
  fmt.Printf("%v", err)
  return
}
fmt.Printf("%s", bt)
```
输出
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://www.douyacun.com/</loc>
    <lastmod>2020-04-19T17:28:33+08:00</lastmod>
    <changefreq>daily</changefreq>
    <priority>1</priority>
  </url>
</urlset>
```

# Features
- [x]  [Image sitemap 图片](#image-sitemap)
- [x]  [News sitemap 新闻](#news-sitemap)
- [x]  [Video sitemap 视频](#video-sitemap)
- [x]  [File storage 文件存储](#file-storage)
- [x]  [Sitemap index](#sitemap-index)

# Image sitemap

Google 图片扩展功能 [Google Image Support](https://support.google.com/webmasters/answer/178636?hl=zh-Hans&ref_topic=4581190)

- **提供适当的相关信息**：确保您的视觉内容与其所在网页的主题相关。我们建议您仅在能为网页增添原创价值的情况下展示图片。我们极不赞成在网页中完全使用非原创的图片和文字内容。
- **优化放置位置**：尽可能将图片放置在相关文字附近。必要时，也可考虑将最重要的图片放置在网页顶部附近。
- **勿将重要文字内嵌在图片中**：避免将文字（特别是网页标题和菜单项等重要的文字元素）内嵌在图片中，因为并非所有用户都能访问这类文字（而且网页翻译工具不适用于图片）。为了尽可能让更多的人能访问您的内容，请使用 HTML 格式提供文本，并为图片提供替代文本。
- **创建信息丰富的优质网站**：对 Google 图片而言，优质的网页内容与视觉内容同等重要 - 它可以提供背景信息并更能吸引用户点击搜索结果。网页内容可用于为图片生成一段文本摘要，而且 Google 在进行图片排名时会考虑[对应的网页内容质量](https://webmasters.googleblog.com/2011/05/more-guidance-on-building-high-quality.html)。
- **创建适合在各种设备上访问的网站**：比起桌面设备，用户更多地使用移动设备在 Google 图片上进行搜索。因此，有必要[设计一个适合所有设备类型和尺寸的网站](https://developers.google.com/search/mobile-sites/)。请使用[移动设备适合性测试工具](https://search.google.com/test/mobile-friendly)测试您的网页在移动设备上的运行效果，并获取反馈以了解哪些内容需要修正。
- **为图片创建良好的网址结构**：Google 会借助网址路径以及文件名来理解您的图片。我们建议您好好组织图片内容，以使网址结构合乎逻辑。

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

Google 新闻站点地图准则 [google news support](https://support.google.com/webmasters/answer/178636?hl=zh-Hans&ref_topic=4581190) 

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

视频 Sitemap 及其替代方案 [Google Video Support](https://support.google.com/webmasters/answer/80471?hl=zh-Hans&ref_topic=4581190)

- **Google 必须能够找到该视频。** 系统会根据是否存在某种 HTML 标记（例如 ``、`` 或 ``）来识别网页中的视频。请确保相应网页不[需要复杂的用户操作或特定的网址片段即可加载](https://support.google.com/webmasters/answer/156442#complex-javascript)，否则 Google 可能找不到它。**提示**：虽然我们可以通过自然抓取找到网页中内嵌的视频，但您也可以通过发布[视频 Sitemap](https://support.google.com/webmasters/answer/80471) 帮助我们找到您的视频。
- **您必须[为视频提供高品质的缩略图](https://support.google.com/webmasters/answer/156442#thumbnails)。**
- **确保每个视频都位于可公开访问的网页中**，用户可以在其中观看视频。该网页不应该要求用户登录，也不应该被 [robots.txt](https://support.google.com/webmasters/answer/6062608) 或 [noindex](https://support.google.com/webmasters/answer/93710) 屏蔽（必须可供 Google 访问）。
- **视频内容应确切吻合其托管网页的内容。** 例如，如果您拥有一个介绍桃饼的食谱网页，不要嵌入笼统介绍甜点的视频。
- 确保您在视频 Sitemap 或视频标记中提供的任何信息与实际视频内容**一致**。

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

# file storage

```go
st := NewSiteMap()
// 默认域名
st.SetDefaultHost("https://www.douyacun.com")
st.SetPretty(true)
// 每个sitemap文件不能多于50000个链接，这里可以自己配置每个文件最多，如果超过MaxLinks，会自动生成sitemap_index.xml文件
// st.SetMaxLinks(2)
st.SetPublicPath("/tmp/gositemap")

url := NewUrl()
url.SetLoc("https://www.douyacun.com/business/article55.html")
url.AppendNews(NewNews().SetName("《示例时报》").
               SetTitle("公司 A 和 B 正在进行合并谈判").
               SetLanguage("zh-cn").
               SetPublicationDate(time.Now()))
st.AppendUrl(url)

path, err := st.Storage()
if err != nil {
  fmt.Printf("%v", err)
  return
}
fmt.Println(path)
```

生成`sitemap.xml`文件

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
      <news:publication_date>2020-04-19T17:44:33+08:00</news:publication_date>
      <news:title>公司 A 和 B 正在进行合并谈判</news:title>
    </news:news>
  </url>
</urlset>
```

# Sitemap index 

拆分较大的站点地图

```go
mapIndex := NewSiteMapIndex()

st1 := NewSiteMap()
st1.SetPretty(true)
st1.SetFilename("sitemap1")
st1.SetCompress(true)

# 以news为例
url := NewUrl()
url.SetLoc("https://www.douyacun.com/business/article55.html")
url.AppendNews(NewNews().SetName("《示例时报》").
               SetTitle("公司 A 和 B 正在进行合并谈判").
               SetLanguage("zh-cn").
               SetPublicationDate(time.Now()))
st1.AppendUrl(url)

url2 := NewUrl().SetLoc("https://www.douyacun.com/business/article56.html")
url2.AppendNews(NewNews().SetName("《示例时报》").
               SetTitle("公司 A 和 C 正在进行合并谈判").
               SetLanguage("zh-cn").
               SetPublicationDate(time.Now()))
st1.AppendUrl(url2)
st1Filename, err := st1.Storage()
if err != nil {
  fmt.Printf("%v", err)
  return
}
mapIndex.Append("https://www.douyacun.com/" + st1Filename)
filename, err := mapIndex.Storage("/tmp/sitemap_index.xml")
if err != nil {
  fmt.Printf("%v", err)
  return
}
fmt.Printf("%v", filename)
```

使用sitemap_index时，建议每个单独的sietmap comporess压缩成.gz文件，`SetCompress`后会自动添加 `.gz`后缀名 ,  生成`sitemap1.xml.gz` 和 `sitemap_index.xml`


