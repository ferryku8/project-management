package seed

import (
	"log"

	"github.com/ferryku8/project-management/config"
	"github.com/ferryku8/project-management/models"
	"github.com/ferryku8/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Panicln("Admin user seeded")
	}
}
