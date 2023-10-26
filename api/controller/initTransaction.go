package controller

import (
	"html/template"
	"net/http"
	"webpaygo/api/models"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
)

// iniciar transaccion
func InitTransaction(w http.ResponseWriter, r *http.Request) {

	var orderId string
	var sessionID string
	var monto int
	var urlRetorno string

	for _, data := range models.DatoTransaction {
		orderId = data.OrdenID
		sessionID = data.SessionID
		monto = data.Monto
		urlRetorno = data.UrlRetorno

	}

	transaction, err := transaction.Create(orderId, sessionID, monto, urlRetorno)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	view := template.Must(template.ParseGlob("views/*"))

	data := map[string]interface{}{
		"url":   transaction.URL,
		"token": transaction.Token,
	}

	err = view.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}
