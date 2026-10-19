package main
import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/database/mongodb"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
      ctx := context.Background()
	database, err := mongodb.InitConnection(ctx)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	userController := initDependencies(database)

	router := gin.Default()

	// Rotas
	router.POST("/user", userController.CreateUser)
	router.GET("/getUserById/:userId", userController.FindUserByID)
	router.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	router.PUT("/user/:userId", userController.UpdateUser)
	router.DELETE("/user/:userId", userController.DeleteUser)

     port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
