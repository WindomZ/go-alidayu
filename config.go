package alidayu

const (
	URL_PRODUCTION string = "https://eco.taobao.com/router/rest"      // 正式环境 (或者: http://gw.api.taobao.com/router/rest)
	URL_SANDBOX    string = "http://gw.api.tbsandbox.com/router/rest" // 沙箱环境 (或者: http://gw.api.tbsandbox.com/router/rest)
)

const (
	DEFAULT_COURIER_SIZE  int = 1  // 默认线程数
	DEFAULT_CAPACITY_SIZE int = 30 // 默认线程承载数量
)

var (
	AppURL    string = URL_SANDBOX // alidayu的api地址
	AppKey    string               // alidayu的key
	AppSecret string               // alidayu的secret
)
