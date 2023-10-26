package main

import (
	"fmt"
	"webpaygo/api/handler"
	"webpaygo/api/models"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)

func main() {

	// Accede a la variable DatoTransaction desde el paquete models
	transactions := models.DatoTransaction

	// Imprime el contenido de DatoTransaction en la consola
	fmt.Println("Contenido de DatoTransaction:")
	for _, transaction := range transactions {
		fmt.Printf("OrdenID: %s, SessionID: %s, Monto: %d, UrlRetorno: %s\n",
			transaction.OrdenID, transaction.SessionID, transaction.Monto, transaction.UrlRetorno)
	}

	webpayplus.SetEnvironmentIntegration()

	handler.Init()

}
