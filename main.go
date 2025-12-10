package main

import (
	"fmt"
	"os"
	"todo-list/internal/app"
	"todo-list/internal/core"
	"todo-list/internal/presenter"
	"todo-list/internal/service"
	"todo-list/internal/storage"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// initialize storage
	jsonStorage, err := storage.GetStorage()
	if err != nil {
		return err
	}

	// initalize tracker
	tracker := core.GetTaskTracker()

	// load the existing file
	if err := jsonStorage.Load(tracker); err != nil {
		return err
	}

	// initalize the service layer
	service := service.NewTaskService(tracker, jsonStorage)

	// initalize the presenter
	presenter := presenter.NewConsolePresenter(service)

	// run the app
	app := app.NewApp(service, presenter)
	if err := app.Run(os.Args[1:]); err != nil {
		return err
	}

	return nil
}
