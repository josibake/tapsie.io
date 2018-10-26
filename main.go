package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
// this is a fake function for now
// will be replaced with something that queries the database
func getMarkers(amenities []string) []Marker {
	// mock data for testing
	var markers []Marker
	if amenities[0] == "dogs" {
		markers = []Marker{
			{Lat: 47.608013, Long: -122.335167},
		}
	}
	if len(amenities) > 1 && amenities[1] == "kids" {
		markers = []Marker{
			{Lat: 37.5483, Long: -121.9886},
			{Lat: 47.608013, Long: -122.335167},
		}
	}
	return markers
}

// return markers when a new request comes in
func markerHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["amenities"]
	if ok {
		fmt.Println(keys)
		markers := getMarkers(keys)
		json.NewEncoder(w).Encode(markers)
	} else {
		// return all markers
		markers := getMarkers([]string{"dogs", "kids"})
		json.NewEncoder(w).Encode(markers)
	}
}

func main() {
	// get port on app startup
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s...", port[1:])

	// create the router
	router := mux.NewRouter().StrictSlash(true)

	// default
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})
	router.HandleFunc("/breweries", markerHandler)
	http.ListenAndServe(port, router)
}
