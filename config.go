package alidayu

const (
	URL_PRODUCTION string = "https://eco.taobao.com/router/rest"      // 正式环境 (或者: http://gw.api.taobao.com/router/rest)
	URL_SANDBOX    string = "http://gw.api.tbsandbox.com/router/rest" // 沙箱环境 (或者: http://gw.api.tbsandbox.com/router/rest)
)

const (
	DEFAULT_COURIER_SIZE  int = 3  // 默认线程数
	DEFAULT_CAPACITY_SIZE int = 10 // 默认线程承载数量
)

var (
	AppURL    string = URL_SANDBOX
	AppKey    string
	AppSecret string
)
