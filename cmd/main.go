package main

import (
	"Inter/webhook/pkg/handler"
	model "Inter/webhook/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)
var wg sync.WaitGroup

func main() {
	requestChannel := make(chan model.RequestData, 10) // Adjust buffer size based on your needs
	var wg sync.WaitGroup

	// Start worker
	go handler.Worker(requestChannel, &wg)

	// Start HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var requestData model.RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send the request to the channel
		wg.Add(1)
		requestChannel <- requestData

		// Respond to the client
		fmt.Fprintf(w, "Request received and being processed.")
	})

	// Start the server
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// Close the channel when the server exits
	close(requestChannel)

	// Wait for the worker to finish processing
	wg.Wait()
}