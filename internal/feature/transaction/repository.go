package transaction

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	primaryDB *mongo.Database
}

func NewRepository(primaryDB *mongo.Client) *Repository {
	db := primaryDB.Database("transaction")
	return &Repository{
		primaryDB: db,
	}
}

func (r *Repository) AddTransaction(ctx context.Context, userid uuid.UUID, amount float32, typetransaction string) error {
	collection := r.primaryDB.Collection("transactions")

	transaction := Transaction{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		WalletId:  userid.String(),
		Amount:    amount,
		Type:      typetransaction,
	}

	_, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	return nil
}
