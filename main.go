package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version   = getEnv("VERSION", "v1.0.0")
	startTime = time.Now()
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	Uptime    string    `json:"uptime"`
	Timestamp time.Time `json:"timestamp"`
}

type InfoResponse struct {
	AppName   string `json:"app_name"`
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	GoVersion string `json:"go_version"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Round(time.Second)
	response := HealthResponse{
		Status:    "healthy",
		Version:   version,
		Uptime:    uptime.String(),
		Timestamp: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		AppName:   "minusone-demo",
		Version:   version,
		BuildTime: time.Now().Format(time.RFC3339),
		GoVersion: "go1.21",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MinusOne Demo Application - Version %s\n", version)
	fmt.Fprintf(w, "Uptime: %s\n", time.Since(startTime).Round(time.Second))
}

func main() {
	port := getEnv("PORT", "8080")
	
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)
	
	log.Printf("Starting MinusOne Demo Server v%s on port %s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
