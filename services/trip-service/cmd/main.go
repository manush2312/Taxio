package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
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

	t, err := svc.CreateTrip(ctx, fare) // creating a trip with the given fare
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// keeping the service alive for demonstration purposes
	for {
		time.Sleep(time.Second)
	}
}
