package account

import (
	"VERITRAN/internal/domain/account/entities"
	"VERITRAN/internal/platform/storage/database/driver"
	"VERITRAN/internal/platform/storage/database/model"
	"errors"
)

type AccountRepository struct {
	driverSql driver.DriverSql
}

func NewAccountRepository(driverSql *driver.DriverSql) *AccountRepository {
	return &AccountRepository{
		driverSql: *driverSql,
	}
}

func (a *AccountRepository) Deposit(nickName string, account *entities.Account) (interface{}, error) {
	modelRecord := model.NewAccountModel(*account.GetNickName(), *account.Getname(), *account.GetLastName(), *account.GetType(), float32(*account.GetAmount()), *account.GetActive())

	result := a.driverSql.Update(nickName, modelRecord)
	if !result {
		return 0, errors.New("error doing deposit")
	}
	return modelRecord, nil
}

func (a *AccountRepository) Get(nickName string) (*map[string]interface{}, error) {
	account := a.driverSql.Get(nickName)
	return account, nil

}

func (a *AccountRepository) WithDrawal(nickName string, account *entities.Account) (float32, error) {
	modelRecord := model.NewAccountModel(*account.GetNickName(), *account.Getname(), *account.GetLastName(), *account.GetType(), float32(*account.GetAmount()), *account.GetActive())

	result := a.driverSql.Update(nickName, modelRecord)
	if !result {
		return 0, errors.New("error doing withdrawal")
	}
	return modelRecord.Amount, nil
}

func (a *AccountRepository) Create(account *entities.Account) error {
	modelRecord := model.NewAccountModel(*account.GetNickName(), *account.Getname(), *account.GetLastName(), *account.GetType(), float32(*account.GetAmount()), *account.GetActive())
	result := a.driverSql.Create(*account.GetNickName(), modelRecord)
	if !result {
		return errors.New("ERROR CREATE OBJECT SQL")
	}
	return nil
}
