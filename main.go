package main

import (
	"log"
	_ "payme/docs"
	"payme/config"
	"payme/internal/api"
	v1 "payme/internal/api/v1"
	"payme/internal/api/v1/task"
	"payme/internal/api/v1/user"
	"payme/internal/repository"
	"payme/internal/services"
	"payme/pkg/database"
)

//	@title			payme API
//	@version		0.0
//	@description	This is the backend API for codetest.
//	@termsOfService	http://myapp.com/terms/

//	@contact.name	NBA Yeabsira
//	@contact.url	http://payme.com/support
//	@contact.email	support@myapp.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host		localhost:8080
// @BasePath	/
func main() {
	dbClient, err := database.ConnectDatabase(config.CFG.MONGO_URI)
	if err != nil {
		log.Fatalf("ERROR: failed to connect database %v", err.Error())
	}

	server := api.NewGinServer()
	database := dbClient.Database(config.CFG.MONGO_DATABASE)

	// user dependencies
	userCollection := database.Collection("user")
	userRepository := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
	userRouteGroup := server.NewRouteSubGroup("v1", "users")
	userHandler := user.NewUserHandler(userService)
	userRoutes := user.NewUserRoutes(userHandler)
	userRoutes.RegisterRoutes(userRouteGroup)

	// auth dependencies
	authRouteGroup := server.NewRouteSubGroup("v1", "auth")
	authService := services.NewAuthServices(userRepository)
	authHandler := v1.NewAuthHandler(authService, userService)
	authRoutes := v1.NewAuthRoutes(authHandler)
	authRoutes.RegisterRoutes(authRouteGroup)

	// task dependencies
	taskCollection := database.Collection("task")
	taskRepository := repository.NewTaskRepository(taskCollection)
	taskRouteGroup := server.NewRouteSubGroup("v1", "tasks")
	taskService := services.NewTaskService(taskRepository)
	taskHandler := task.NewTaskHandler(taskService)
	taskRoutes := task.NewTaskRoutes(taskHandler)
	taskRoutes.RegisterRoutes(taskRouteGroup)

	log.Printf("INFO: starting the server at address %v\n", config.CFG.ADDR)
	if err := server.Run(config.CFG.ADDR); err != nil {
		log.Fatalln("ERROR: failed to start the server " + err.Error())
	}
}
