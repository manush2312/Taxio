package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	// Entry point of the trip-service application

	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository() // using in-memory repository for now, or creating a temporary DB
	svc := service.NewService(inmemRepo)         // creating a service layer and injecting the in-memory repository

	fare := &domain.RideFareModel{
		UserID:            "42",
		PackageSlug:       "sedan",
		TotalPriceInCents: 1500.0,
	}

	svc.CreateTrip(ctx, fare) // creating a trip with the given fare
}
