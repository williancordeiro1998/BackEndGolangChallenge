package main

import (
	"log"
	"net/http"
	"path/filepath"

	"backendgolangchallenge/pkg/config"
	"backendgolangchallenge/pkg/controller"
	"backendgolangchallenge/pkg/process"

	"github.com/gorilla/mux"
)

func main() {
	config.Init()

	router := mux.NewRouter()

	// Configurar manipulador para arquivos estáticos (HTML, CSS, JS)
	staticDir := "/static/"
	staticPath := filepath.Join("C:/Program Files/Go/src/backendgolangchallenge", "static")
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(staticPath))))

	bitcoinController := &controller.BitcoinController{}

	router.HandleFunc("/details/{address}", bitcoinController.BitcoinDetailsHandler).Methods("GET")
	router.HandleFunc("/balance/{address}", bitcoinController.BitcoinBalanceHandler).Methods("GET")
	router.HandleFunc("/send", bitcoinController.BitcoinSendHandler).Methods("POST")
	router.HandleFunc("/tx/{tx}", bitcoinController.GetTransactionDetails).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	router.HandleFunc("/process", process.ProcessData).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Servidor iniciado em http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
