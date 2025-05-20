package main

import (
	"context"
	"log"

	"github.com/ale-neto/golang/src/config/database/mongodb"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
