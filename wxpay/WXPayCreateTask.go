package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type WXPayJSAPIData struct {
	AppId     string `json:"appId"`
	TimeStamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type WXPayCreateTaskResult struct {
	app.Result
	Data interface{} `json:"data,omitempty"`
}

const WXPayTradeTypeJSAPI = "JSAPI"
const WXPayTradeTypeNATIVE = "NATIVE"
const WXPayTradeTypeAPP = "APP"

type WXPayCreateTask struct {
	app.Task
	TradeId   string `json:"tradeId"`
	TradeType string `json:"tradeType"`
	NonceStr  string `json:"nonceStr"`
	Openid    string `json:"openid"`
	Value     int64  `json:"value"`
	Body      string `json:"body"`
	ClientIp  string `json:"clientIp"`
	Result    WXPayCreateTaskResult
}

func (task *WXPayCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *WXPayCreateTask) GetInhertType() string {
	return "wxpay"
}

func (task *WXPayCreateTask) GetClientName() string {
	return "WXPay.Create"
}
