package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webpaygo/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {

	router := mux.NewRouter()

	//view

	router.HandleFunc("/", controller.InitTransaction).Methods("GET")
	router.HandleFunc("/commit", controller.VerifTransaction).Methods("POST")
	router.HandleFunc("/save-transaction", controller.SaveTransaction).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	fmt.Println("server en escucha")

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
