package mongo

import (
	"context"
	"time"

	"github.com/fxh111111/utility/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCli struct {
	client *mongo.Client
	logger logger.Logger
}

func Set(uri string, logger logger.Logger) {
	connect()
}

func connect() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Default.Error(ctx, "mongo connection err", err)
	}
}

func GetMongoClient() *mongo.Client {
	return client
}
