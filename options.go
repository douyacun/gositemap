package gositemap

type options struct {
	defaultHost string
	publicPath  string
	filename    string
	compress    bool
	pretty      bool
}

func NewOptions() *options {
	return &options{
		defaultHost: "http://www.example.com",
		publicPath:  "",
		filename:    "sitemap",
		compress:    true,
		pretty:      false,
	}
}

func (o *options) SetDefaultHost(host string) {
	o.defaultHost = host
}

func (o *options) SetPublicPath(path string) {
	o.publicPath = path
}

func (o *options) SetFilename(filename string) {
	o.filename = filename
}

func (o *options) SetCompress(compress bool) {
	o.compress = compress
}

func (o *options) SetPretty(pretty bool) {
	o.pretty = pretty
}
