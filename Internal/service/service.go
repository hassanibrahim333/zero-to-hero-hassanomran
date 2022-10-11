package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"secondchallange/Internal/Models"
	"secondchallange/Internal/Repository"
)

type DefaultTransactionService struct {
	transactionRepo Repository.IRepo
}

func NewDefaultService(repo Repository.IRepo) *DefaultTransactionService {
	return &DefaultTransactionService{
		transactionRepo: repo,
	}
}
func (s *DefaultTransactionService) List(ctx context.Context) ([]Models.Transaction, error) {
	models, err := s.transactionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var temp Models.Transaction
	var result []Models.Transaction
	for _, e := range models {
		temp.ID = e.ID
		temp.Amount = e.Amount
		temp.Currency = e.Currency
		temp.CreatedAt = e.CreatedAt
		temp.StatusId = e.StatusId
		temp.AuthorizationId = e.AuthorizationId
		result = append(result, temp)
	}
	return result, nil
}

func (s *DefaultTransactionService) Create(ctx context.Context, transaction *Models.Transaction) (*Models.Transaction, error) {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return nil, err
	}
	model := Repository.Transactions{
		ID:              transaction.ID,
		Amount:          transaction.Amount,
		Currency:        transaction.Currency,
		CreatedAt:       transaction.CreatedAt,
		StatusId:        transaction.StatusId,
		AuthorizationId: transaction.AuthorizationId,
	}
	if err := s.transactionRepo.Create(ctx, &model); err != nil {
		return nil, err
	}
	res := Models.Transaction{
		ID:              model.ID,
		Amount:          model.Amount,
		Currency:        model.Currency,
		CreatedAt:       model.CreatedAt,
		StatusId:        model.StatusId,
		AuthorizationId: model.AuthorizationId,
	}
	return &res, nil
}
