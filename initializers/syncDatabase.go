package initializers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Tag{},
		&models.ShortUrl{},
	)
	if err != nil {
		fmt.Println(err)
	}
}
