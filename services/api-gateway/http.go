package main

// ! This file contains the actual function that handles HTTP requests.

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {

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

	// TODO : Call the trip service (gRPC call)

	response := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, response)
}
