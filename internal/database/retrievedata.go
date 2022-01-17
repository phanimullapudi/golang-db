package database

import (
	"fmt"

	"github.com/phanimullapudi/golang-db/cmd/user-processing"
	"gorm.io/gorm"
)

var users *user.User

func RetriveData(db *gorm.DB) error {

	fmt.Println("Retrieving Data")
	db.First(&users)
	result := db.Find(&users, 10)
	fmt.Println(result.RowsAffected)
	return nil

}
