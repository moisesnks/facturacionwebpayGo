package main

import (
	"webpaygo/handler"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)

func main() {

	webpayplus.SetEnvironmentIntegration()

	handler.Init()

}
