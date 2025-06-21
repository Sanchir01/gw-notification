package transaction

import (
	"context"
	"github.com/google/uuid"
)

type TransactionService interface {
	AddTransaction(ctx context.Context, userid uuid.UUID, amount float32, typetransaction string) error
}
type Service struct {
	repo TransactionService
}

func NewService(repo TransactionService) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SetTransaction(ctx context.Context, userid uuid.UUID, amount float32, typetransaction string) error {
	if err := s.repo.AddTransaction(ctx, userid, amount, typetransaction); err != nil {
		return err
	}
	return nil
}
