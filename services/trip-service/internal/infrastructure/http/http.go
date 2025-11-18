package http

// this file contains HTTP handlers for trip-related operations and basically it is a transport layer.

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/types"
)

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var requestBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

}
