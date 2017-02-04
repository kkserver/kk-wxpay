package wxpay

import (
	"crypto/tls"
	"crypto/x509"
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
	GetCerts() ([]tls.Certificate, error)
	GetCA() *x509.CertPool
}

type WXPayApp struct {
	app.App

	Remote *remote.Service

	AppId     string
	Key       string
	MchId     string
	Prefix    string
	NotifyUrl string

	Cert    string
	CertKey string

	WXPay *WXPayService

	ca    *x509.CertPool
	certs []tls.Certificate
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

func (C *WXPayApp) GetCA() *x509.CertPool {
	if C.ca == nil {
		C.ca = x509.NewCertPool()
		C.ca.AppendCertsFromPEM(pemCerts)
	}
	return C.ca
}

func (C *WXPayApp) GetCerts() ([]tls.Certificate, error) {

	if C.certs != nil {
		return C.certs, nil
	}

	cliCrt, err := tls.LoadX509KeyPair(C.Cert, C.CertKey)

	if err != nil {
		return nil, err
	}

	C.certs = []tls.Certificate{cliCrt}

	return C.certs, nil
}
