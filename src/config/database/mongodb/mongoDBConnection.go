package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URI     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

// NewMongoDBConnection estabelece conexão com MongoDB
// Abordagem usando reflection/interface{} para tentar diferentes possibilidades
func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URI)

	mongodb_database := os.Getenv(MONGODB_USER_DB)

	// Cria as opções do cliente
	clientOptions := options.Client().ApplyURI(mongodb_uri)

	// Tenta diferentes abordagens
	var client *mongo.Client
	var err error

	// Abordagem 1: Tenta com opções primeiro
	client, err = mongo.Connect(clientOptions)

	// Se não funcionar, tente outras abordagens

	if client == nil || err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
