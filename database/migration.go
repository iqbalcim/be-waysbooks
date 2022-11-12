package database

import (
	"fmt"
	"waysbooks/models"
	"waysbooks/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Cart{},
	)
	
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}