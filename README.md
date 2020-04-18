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

# Features
- image 图片
- news 新闻
