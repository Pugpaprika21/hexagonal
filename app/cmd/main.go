package main

import (
	"log"
	"os"
	"os/signal"
	"pugpaprika/app/internal/adapter/handler"
	"pugpaprika/app/internal/adapter/repository"
	"pugpaprika/app/internal/adapter/router"
	"pugpaprika/app/internal/core/service"
	"pugpaprika/app/pkg/config"
	"pugpaprika/app/pkg/factory"
	"syscall"
)

func main() {
	if err := config.Loadenvironment(); err != nil {
		log.Fatalf("Error loading environment config: %s", err)
	}

	connector, err := factory.NewDatabase()
	if err != nil {
		log.Fatalf("Error selecting database: %s", err)
	}

	db, err := connector.Connect()
	if err != nil {
		log.Fatalf("Error establishing database connection: %s", err)
	}

	defer func() {
		if err := connector.Close(); err != nil {
			log.Printf("Error closing database connection: %s", err)
		}
	}()

	repos := repository.NewRepository(db)
	servs := service.NewService(repos)
	handlers := handler.NewHandler(servs)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		log.Printf("Received signal: %s. Shutting down...", sig)
		if err := connector.Close(); err != nil {
			log.Printf("Error closing database connection during shutdown: %s", err)
		}

		os.Exit(0)
	}()

	router.NewRouter(handlers).Start()
}
