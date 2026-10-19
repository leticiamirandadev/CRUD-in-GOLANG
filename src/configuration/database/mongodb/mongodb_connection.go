package mongodb
import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func InitConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	mongoDatabase := os.Getenv("MONGODB_DB")
	if mongoDatabase == "" {
		mongoDatabase = "crud-golang"
	}
      client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao MongoDB com sucesso!")
      return client.Database(mongoDatabase), nil
}
