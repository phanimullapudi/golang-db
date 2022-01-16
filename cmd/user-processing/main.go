package main

import (
	"fmt"

	"github.com/phanimullapudi/golang-db/internal/database"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	fmt.Println("Setting up the App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	fmt.Println("Connected to the database")

}
