package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// RideFareModel and TripModel --> describe business data that our app reasons about.
type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // type of vehicle that we are selecting --> ex. van, sedan etc
	TotalPriceInCents float64
}
