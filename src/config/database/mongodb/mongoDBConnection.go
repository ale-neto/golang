package mongodb

import (
	"context"
	"time"

	"github.com/ale-neto/golang/src/config/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitConnection() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Ping(ctx, nil); err != nil {
			panic(err)
		}
	}

	logger.Info("MongoDB connection established successfully")
}
