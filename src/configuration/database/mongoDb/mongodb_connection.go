package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

// Construtor ou variavel global?
// De preferencia ao projeto, se vc esta usando um construitor em tudo (que é o caso)
// Continue utilizando construtores
// var (Connection mongo.Client)
func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	//logger.Info("Conseguiu conectar!")
	return client.Database(mongodb_database), nil

}

//Usado apenas para testar a conexão
// func InitConnection() {
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:root@localhost:27017/?authSource=admin"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := client.Ping(ctx, nil); err != nil {
// 		panic(err)
// 	}

// 	logger.Info("Conseguiu conectar!")

// }
