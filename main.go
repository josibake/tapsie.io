package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// lil marker
type Marker struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// get the port from the ENV on startup
func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// get markers based on an id
func getMarkers(id int) []Marker {
	// mock data for testing
	markers := []Marker{
		{Lat: 47.608013, Long: -122.335167},
		{Lat: 37.5483, Long: -121.9886},
	}
	return markers
}

// return markers when a new request comes in
func markerHandler(w http.ResponseWriter, r *http.Request) {
	markers := getMarkers(123)
	json.NewEncoder(w).Encode(markers)
}

func main() {
	// get port on app startup
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s...", port[1:])

	// return an array of markers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})
	http.HandleFunc("/breweries", markerHandler)
	http.ListenAndServe(port, nil)
}
