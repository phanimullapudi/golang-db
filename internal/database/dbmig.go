package database

import (
	"fmt"

	"github.com/phanimullapudi/golang-db/cmd/user-processing"
	"gorm.io/gorm"
)

// MigrateDB - function to sync the databse with User struct
func MigrateDB(db *gorm.DB) error {
	fmt.Println("Conducting DB migration")
	db.AutoMigrate(&user.User{})

	return nil

}
