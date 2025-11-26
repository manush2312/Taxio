package http

// this file contains HTTP handlers for trip-related operations and basically it is a transport layer.

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

// A struct that holds TripService.
type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var requestBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// fare := &domain.RideFareModel{
	// 	UserID:            "42",
	// 	PackageSlug:       "sedan",
	// 	TotalPriceInCents: 1500.0,
	// }

	ctx := r.Context()

	t, err := s.Service.GetRoute(ctx, &requestBody.Pickup, &requestBody.Destination) // creating a trip with the given fare
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusOK, t)

}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) // encoding the data into JSON and writing it to the response writer
}
