package main

import (
	"log"
	"task-api/authmiddleware"
	"task-api/config"
	"task-api/db"
	"task-api/routes"

	redis "task-api/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("C:/Users/DELL/OneDrive/Desktop/task-api/.env")
	if err != nil {
		log.Println("No .env file found")
	}
	// Load the config
	cfg := config.LoadConfig()

	// Connect to the PostgreSQL database
	db := db.ConnectDB(cfg)
	if db == nil {
		log.Fatal("Failed to connect to PostgreSQL.")
	}

	// Connect to Redis
	rdb := redis.ConnectRedis()
	if rdb == nil {
		log.Fatal("Failed to connect to Redis.")
	}

	// Set up Gin router
	r := gin.Default()

	// Routes
	r.POST("/login", routes.Login)
	r.POST("/register", routes.Register)

	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(authmiddleware.AuthMiddleware())
	{
		taskRoutes.POST("create", routes.CreateTask)
		taskRoutes.GET("view", routes.GetTasks)
		taskRoutes.GET("view/id/:id", routes.GetTaskByID)
		taskRoutes.PUT(":id", routes.UpdateTask)
		taskRoutes.DELETE(":id", routes.DeleteTask)
	}

	r.Run(":8080")
}
