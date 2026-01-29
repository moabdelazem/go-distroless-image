package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Config struct {
	SrvPort string
	AppName string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewConfig() *Config {
	return &Config{
		SrvPort: getEnv("SERVER_PORT", ":8080"),
		AppName: getEnv("APP_NAME", "Distroless API"),
	}
}

type Response struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, Response{
		Message: "Hello from a Configurable Distroless API!",
	})
}

func handleHealth(cfg *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, HealthResponse{
			Status:  "healthy",
			Service: cfg.AppName,
		})
	}
}

func main() {
	srvCfg := NewConfig()

	r := mux.NewRouter()
	r.HandleFunc("/", handleRoot).Methods("GET")
	r.HandleFunc("/health", handleHealth(srvCfg)).Methods("GET")

	log.Printf("Server starting on port %s", srvCfg.SrvPort)
	if err := http.ListenAndServe(srvCfg.SrvPort, r); err != nil {
		log.Fatal(err)
	}
}
