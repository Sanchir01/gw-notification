package app

import (
	"github.com/Sanchir01/gw-notification/internal/feature/transaction"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	TransactionRepository *transaction.Repository
}

func NewRepositories(primaryDB *mongo.Client) *Repositories {
	return &Repositories{
		TransactionRepository: transaction.NewRepository(primaryDB),
	}
}
