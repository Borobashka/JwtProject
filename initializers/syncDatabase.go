package initializers

import (
	"jwtProject/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}