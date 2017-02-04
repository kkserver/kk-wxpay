package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type WXRefundTaskResult struct {
	app.Result
	RefundNo string `json:"refundId,omitempty"`
}

type WXRefundTask struct {
	app.Task
	TradeId       string `json:"tradeId"`
	TransactionId string `json:"transactionId"`
	RefundId      string `json:"refundId"`
	Value         int64  `json:"value"`
	RefundValue   int64  `json:"refundValue"`
	Result        WXRefundTaskResult
}

func (task *WXRefundTask) GetResult() interface{} {
	return &task.Result
}

func (task *WXRefundTask) GetInhertType() string {
	return "wxpay"
}

func (task *WXRefundTask) GetClientName() string {
	return "WXPay.Refund"
}
