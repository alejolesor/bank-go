package repository

import "VERITRAN/internal/domain/account/entities"

type AccountRepository interface {
	Deposit(nickName string, account *entities.Account) (interface{}, error)
	WithDrawal(nickName string, account *entities.Account) (float32, error)
	Get(nickName string) (*map[string]interface{}, error)
	Create(account *entities.Account) error
}
