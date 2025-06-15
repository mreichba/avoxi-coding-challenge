package main

import (
	"avoxi-geoip/geoip"
	"avoxi-geoip/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize GeoIP database
	err := geoip.InitDB("Data/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatalf("Error initializing GeoIP database: %v", err)
	}
	defer geoip.CloseDB()

	//Initiate router
	r := mux.NewRouter()

	//this was recomemended by ai as i was having CORS issues and it was not working properly
	// add CORS headers
	r.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
	}).Methods("OPTIONS")

	// register /verify endpoint
	r.HandleFunc(("/verify"), handlers.VerifyIP).Methods("POST")

	// start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
