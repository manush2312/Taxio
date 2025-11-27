package main

// ! This file contains the actual function that handles HTTP requests.

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"time"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 9)

	var reqBody previewTripRequest
	// unmarshal the JSON request body into the reqBody struct, according to the previewTripRequest type defined in types.go
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parese JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validation that userID is present
	if reqBody.UserID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	// TODO : Call the trip service (gRPC call)
	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Print(err)
		return
	}

	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to parese JSON data from trip service", http.StatusBadRequest)
		return
	}
	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusCreated, response)
}
