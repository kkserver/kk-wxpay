package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type WXPayConfirmData struct {
	Openid        string `json:"openid"`
	TradeId       string `json:"tradeId"`
	TransactionId string `json:"transactionId"`
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
