// internal/api/server.go
package api

import (
	"keeper/internal/storage" // Assicurati che il path del modulo sia corretto
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	store      storage.Store
}

func NewAPIServer(listenAddr string, store storage.Store) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// Run avvia il server HTTP e registra tutte le rotte (gli endpoint).
func (s *APIServer) Run() {
	// Creiamo un nuovo router. Sarà il nostro "centralino" per le richieste.
	router := http.NewServeMux()

	// Registriamo l'endpoint di health check.
	// Diciamo al router: "Quando ricevi una richiesta GET a /healthcheck,
	// esegui la funzione handleHealthCheck".
	router.HandleFunc("GET /healthcheck", s.handleHealthCheck)

	log.Println("JSON API server running on port", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}
