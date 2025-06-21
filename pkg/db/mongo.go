package db

import (
	"context"
	"github.com/Sanchir01/gw-notification/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoClient(ctx context.Context, host, port, database string) (client *mongo.Client, err error) {
	uri := "mongodb://" + host + ":" + port + "/" + database
	clientOptions := options.Client().ApplyURI(uri)

	utils.DoWithTries(func() error {
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			return err
		}
		return nil
	}, 5, time.Second*2)

	return client, nil
}
