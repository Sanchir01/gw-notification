package app

import (
	"github.com/Sanchir01/gw-notification/internal/feature/transaction"
)

type Services struct {
	TransactionService *transaction.Service
}

func NewServices(r *Repositories) *Services {
	return &Services{
		TransactionService: transaction.NewService(r.TransactionRepository),
	}
}
