package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WXPayService struct {
	app.Service

	Create  *WXPayCreateTask
	Confirm *WXPayConfirmTask

	ca *x509.CertPool
}

func (S *WXPayService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *WXPayService) HandleInitTask(a app.IApp, task *app.InitTask) error {

	S.ca = x509.NewCertPool()
	S.ca.AppendCertsFromPEM(pemCerts)

	return nil
}

func Sign(data map[string]interface{}, secret string) string {

	b := bytes.NewBuffer(nil)

	keys := []string{}

	for key, _ := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		v := dynamic.StringValue(dynamic.Get(data, key), "")
		if v != "" {
			b.WriteString(key)
			b.WriteString("=")
			b.WriteString(dynamic.StringValue(v, ""))
			b.WriteString("&")
		}
	}

	b.WriteString("key=")
	b.WriteString(secret)

	log.Println(b.String())

	m := md5.New()
	m.Write(b.Bytes())

	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

func NewNonceStr() string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(m.Sum(nil))
}

type WXPayCreateResultData struct {
	ReturnCode string `xml:"return_code",json:"return_code"`
	ReturnMsg  string `xml:"return_msg",json:"return_msg"`
	AppId      string `xml:"appid",json:"appid"`
	MchId      string `xml:"mch_id",json:"mch_id"`
	NonceStr   string `xml:"nonce_str",json:"nonce_str"`
	Openid     string `xml:"openid",json:"openid"`
	Sign       string `xml:"sign",json:"sign"`
	PrepayId   string `xml:"prepay_id",json:"prepay_id"`
	TradeType  string `xml:"trade_type",json:"trade_type"`
}

func (S *WXPayService) HandleWXPayCreateTask(a IWXPayApp, task *WXPayCreateTask) error {

	if task.TradeType == "" {
		task.TradeType = WXPayTradeTypeJSAPI
	}

	if task.NonceStr == "" {
		task.NonceStr = NewNonceStr()
	}

	data := map[string]interface{}{}

	data["appid"] = a.GetAppId()
	data["mch_id"] = a.GetMchId()
	data["nonce_str"] = task.NonceStr
	data["body"] = task.Body
	data["out_trade_no"] = a.GetPrefix() + task.TradeId
	data["total_fee"] = task.Value
	data["spbill_create_ip"] = task.ClientIp
	data["notify_url"] = a.GetNotifyUrl()
	data["trade_type"] = task.TradeType
	data["openid"] = task.Openid
	data["sign_type"] = "MD5"
	data["sign"] = Sign(data, a.GetKey())

	b := bytes.NewBuffer(nil)

	b.WriteString("<xml>")

	for key, value := range data {
		b.WriteString("<")
		b.WriteString(key)
		b.WriteString("><![CDATA[")
		b.WriteString(dynamic.StringValue(value, ""))
		b.WriteString("]]></")
		b.WriteString(key)
		b.WriteString(">")
	}

	b.WriteString("</xml>")

	log.Println(b.String())

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: S.ca},
		},
	}

	resp, err := client.Post("https://api.mch.weixin.qq.com/pay/unifiedorder?", "text/xml; charset=utf-8", b)

	if err != nil {
		task.Result.Errno = ERROR_WXPAY
		task.Result.Errmsg = err.Error()
	} else if resp.StatusCode == 200 {
		var body = make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(body)
		defer resp.Body.Close()

		log.Println(string(body))

		data := WXPayCreateResultData{}

		err = xml.Unmarshal(body, &data)

		if err != nil {
			task.Result.Errno = ERROR_WXPAY
			task.Result.Errmsg = err.Error()
			return nil
		}

		if data.ReturnCode == "SUCCESS" {

			if task.TradeType == WXPayTradeTypeJSAPI {

				pay := map[string]interface{}{}

				pay["appId"] = a.GetAppId()
				pay["timeStamp"] = time.Now().Unix()
				pay["nonceStr"] = task.NonceStr
				pay["package"] = "prepay_id=" + data.PrepayId
				pay["signType"] = "MD5"
				pay["paySign"] = Sign(pay, a.GetKey())
				pay["timestamp"] = pay["timeStamp"]

				delete(pay, "timeStamp")

				task.Result.Data = &pay

			} else {
				task.Result.Data = &data
			}

		} else {
			task.Result.Errno = ERROR_WXPAY
			task.Result.Errmsg = data.ReturnMsg
			return nil
		}

	} else {
		var body = make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(body)
		defer resp.Body.Close()
		log.Println(string(body))
		task.Result.Errno = ERROR_WXPAY
		task.Result.Errmsg = fmt.Sprintf("[%d] %s", resp.StatusCode, string(body))
	}

	return nil
}
