package account

import (
	"VERITRAN/internal/domain/account/entities"
	"VERITRAN/internal/domain/account/repository"
)

type AccountCreate interface {
	Create(account *entities.Account) error
}

type AcountCreateService struct {
	repository repository.AccountRepository
}

func NewCreateUseCase(repository repository.AccountRepository) *AcountCreateService {
	return &AcountCreateService{
		repository: repository,
	}
}

func (a *AcountCreateService) Create(account *entities.Account) error {
	err := a.repository.Create(account)
	if err != nil {
		return err
	}
	return nil
}
