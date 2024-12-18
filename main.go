package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world! Your Go API is working. This is pulled from Github. Again changed. Changed from local."))
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
