package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webpaygo/api/models"
)

// Handler para recibir y guardar la transacción
func SaveTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.DatoTransaction = append(models.DatoTransaction, t)

	fmt.Printf("Received data: %+v\n", t)

	// Opcional: Puedes responder con una confirmación
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Transacción guardada")
}
