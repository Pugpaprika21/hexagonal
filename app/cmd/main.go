package main

import (
	"log"
	"pugpaprika/app/internal/adapter/handler"
	"pugpaprika/app/internal/adapter/repository"
	"pugpaprika/app/internal/adapter/router"
	"pugpaprika/app/internal/core/service"
	"pugpaprika/app/pkg/config"
	"pugpaprika/app/pkg/factory"
)

func main() {
	err := config.Loadenvironment()
	if err != nil {
		log.Fatalf("error loading environment config : %s", err.Error())
	}

	connector, err := factory.NewDatabase()
	if err != nil {
		log.Fatalf("error selecting database: %s", err.Error())
	}

	db, err := connector.Connect()
	if err != nil {
		log.Fatalf("error establishing database connection: %s", err.Error())
	}

	defer func() {
		if err := connector.Close(); err != nil {
			log.Fatalf("error closing database connection: %s", err.Error())
		}
	}()

	repos := repository.NewRepository(db)
	servs := service.NewService(repos)
	handlers := handler.NewHandler(servs)

	router.NewRouter(handlers).Start()
}
