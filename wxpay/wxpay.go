package wxpay

import (
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type IWXPayApp interface {
	app.IApp
	GetAppId() string
	GetKey() string
	GetMchId() string
	GetPrefix() string
	GetNotifyUrl() string
}

type WXPayApp struct {
	app.App

	Remote *remote.Service

	AppId     string
	Key       string
	MchId     string
	Prefix    string
	NotifyUrl string

	WXPay *WXPayService
}

func (A *WXPayApp) GetAppId() string {
	return A.AppId
}

func (A *WXPayApp) GetKey() string {
	return A.Key
}

func (A *WXPayApp) GetMchId() string {
	return A.MchId
}

func (A *WXPayApp) GetPrefix() string {
	return A.Prefix
}

func (A *WXPayApp) GetNotifyUrl() string {
	return A.NotifyUrl
}
