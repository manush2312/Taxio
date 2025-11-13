package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

// this is going to be MongoDB implementation for the repository

// “In-memory repository” means we’re temporarily storing trips in Go memory (maps), not in MongoDB or any external database.
/*
	this is perfect for :-
		1. Early testing or prototyping
		2. Unit testing TripService
	Once our logic is verified then we will add real MongoDB layer later on.

	Basic flow for now : TripService  →  TripRepository interface  →  inmemRepository implementation  →  Go maps (temporary DB)
*/
type inmemRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
