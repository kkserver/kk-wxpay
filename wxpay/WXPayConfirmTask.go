package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type WXPayConfirmData struct {
	TradeId string `json:"tradeId"`
	TradeNo string `json:"tradeNo"`
}

type WXPayConfirmTaskResult struct {
	app.Result
	Data *WXPayConfirmData `json:"data,omitempty"`
}

type WXPayConfirmTask struct {
	app.Task
	Body   string `json:"body"`
	Result WXPayConfirmTaskResult
}

func (task *WXPayConfirmTask) GetResult() interface{} {
	return &task.Result
}

func (task *WXPayConfirmTask) GetInhertType() string {
	return "wxpay"
}

func (task *WXPayConfirmTask) GetClientName() string {
	return "WXPay.Confirm"
}
