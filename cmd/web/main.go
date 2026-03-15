package main

import (
	"log/slog"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	r := app.Setup()

	logger.Info("starting server")

	err := r.Run(":8080")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	os.Exit(1)
}
