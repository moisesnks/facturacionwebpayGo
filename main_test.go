package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"webpaygo/api/controller"
)

func TestVerifTransaction(t *testing.T) {
	// Crea una solicitud HTTP de prueba con el método POST a la ruta "/verifica".
	req, err := http.NewRequest("POST", "/commit", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crea un ResponseRecorder para registrar la respuesta.
	rr := httptest.NewRecorder()

	// Llama a la función que deseas probar pasando la solicitud y el ResponseRecorder.
	controller.VerifTransaction(rr, req)

	// Verifica el código de estado de la respuesta. Si todo está bien, debería ser 200.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}
}
