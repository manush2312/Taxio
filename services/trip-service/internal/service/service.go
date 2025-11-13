package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This file defines your business logic — the “what should happen” part.
/*
It uses the TripRepository (data access) from your domain layer to perform operations like creating a trip, cancelling a trip, etc.
So it sits between our handlers (API layer, HTTP/gRPC) and the repository (data access layer).

Handler → Service → Repository → Database
*/
type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}
