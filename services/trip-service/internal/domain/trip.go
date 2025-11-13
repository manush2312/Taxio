package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID
	UserID   string
	Status   string
	RideFare *RideFareModel
}

// TripRepository knows how to store that trip in MongoDB, This defines what we can do with data — not how.
// it answers the question: "What can we do with Trip data?" or "What storage operations are possible for trips?"
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error) // insert trip into MongoDB/database
}

// This defines your business logic contract — what actions can be performed by the service.
// TripService knows how to create and manage trip
type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}
