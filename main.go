package main

import (
	"log"

	"github.com/ferryku8/project-management/config"
	"github.com/ferryku8/project-management/controllers"
	"github.com/ferryku8/project-management/database/seed"
	"github.com/ferryku8/project-management/repositories"
	"github.com/ferryku8/project-management/routes"
	"github.com/ferryku8/project-management/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserServices(userRepo)
	userController := controllers.NewUserController(userService)

	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))

}
