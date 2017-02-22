package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
)

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
	NotifyUrl string `json:"notifyUrl"`
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
