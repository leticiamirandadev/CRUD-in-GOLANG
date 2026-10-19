package connection

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (*mongo.Database, func()) {
	ctx := context.Background()

	mongodbContainer, err := mongodb.Run(ctx, "mongo:6.0")
	if err != nil {
		panic(err)
	}

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	if err != nil {
		panic(err)
	}

	db := client.Database("test_db")

	closeConnection := func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
		if err := testcontainers.TerminateContainer(mongodbContainer); err != nil {
			panic(err)
		}
	}

	return db, closeConnection
}
