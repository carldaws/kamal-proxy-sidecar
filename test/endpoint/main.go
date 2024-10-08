package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type RequestData struct {
	Body json.RawMessage `json:"body"`
}

var (
	lastRequestData RequestData
	mutex           sync.Mutex
)

func upHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	lastRequestData.Body = body
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Request received")
}

func getLastRequestHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if lastRequestData.Body == nil {
		http.Error(w, "No requests received yet", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(lastRequestData.Body)
}

func main() {
	http.HandleFunc("/up", upHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/last", getLastRequestHandler)

	fmt.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
