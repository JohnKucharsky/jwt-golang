package initializers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	}
}
