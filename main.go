package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	// get port on app startup
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s...", port[1:])

	//hello world example to get started
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var response string
		if r.URL.Path == "/" {
			response = "hello, world!"
		} else {
			response = r.URL.Path[1:]
		}
		fmt.Fprintf(w, "tapsie says: %s\n", response)
	})

	http.ListenAndServe(port, nil)
}
