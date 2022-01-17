package main

import (
	"fmt"

	"github.com/phanimullapudi/golang-db/internal/database"
)

// App - the struct which contains pointers to database connections and etc....
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up the App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	fmt.Println("Connected to the database", db)

	err = database.MigrateDB(db)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		return err
	}

	err = database.RetriveData(db)
	if err != nil {
		return err
	}

	return nil

}

func main() {
	app := App{
		Name:    "Users Processing",
		Version: "1.0",
	}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting the APP", err)
	}
}
