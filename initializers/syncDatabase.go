package initializers

import "github.com/davi-OD/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
