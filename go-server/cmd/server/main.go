package main

import (
	"avoxi-geoip/config"
	"avoxi-geoip/geoip"
	"avoxi-geoip/handlers"
	"avoxi-geoip/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Load Config Variables from .env
	cfg := config.LoadConfig()

	// Initialize GeoIP database with path from config
	if err := geoip.InitDB(cfg.GeoIPDBPath); err != nil {
		log.Fatalf("Error initializing GeoIP database: %v", err)
	}
	defer geoip.CloseDB()

	//Initiate router
	r := mux.NewRouter()

	//Apply CORS middleware
	r.Use(middleware.CORSMiddleware)

	// register /verify endpoint
	r.HandleFunc(("/verify"), handlers.VerifyIP).Methods("POST", "OPTIONS")

	// start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
