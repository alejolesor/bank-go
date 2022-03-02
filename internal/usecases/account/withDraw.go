package account

import (
	"VERITRAN/internal/domain/account/entities"
	"VERITRAN/internal/domain/account/repository"
	serviceDomain "VERITRAN/internal/domain/account/services"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type AccountWithDraw interface {
	WithDrawall(nickName string, withdrawal float64) (float64, error)
}

type AcountWithDrawalService struct {
	serviceDomain serviceDomain.WithDrawService
	repository    repository.AccountRepository
}

func NewWithDrawUseCase(serviceDomain serviceDomain.WithDrawService, repository repository.AccountRepository) *AcountWithDrawalService {
	return &AcountWithDrawalService{
		serviceDomain: serviceDomain,
		repository:    repository,
	}
}

func (a *AcountWithDrawalService) WithDrawall(nickName string, withdrawal float64) (float64, error) {
	account, err := a.getAccount(nickName)
	if err != nil {
		return 0, err
	}
	newBalance, err := a.serviceDomain.ValidWithDraw(*account.GetAmount(), withdrawal)
	if err != nil {
		return 0, nil
	}

	accountCurrently := a.assignNewBalance(account, newBalance)

	currentBalane, err := a.repository.WithDrawal(nickName, accountCurrently)
	if err != nil {
		return 0, err
	}

	return float64(currentBalane), nil
}

func (a *AcountWithDrawalService) getAccount(nickName string) (*entities.Account, error) {
	account, err := a.repository.Get(nickName)
	if err != nil {
		return nil, err
	}
	accountUser := a.convertObject(nickName, *account)
	if *accountUser.GetNickName() == "" {
		return nil, errors.New("account not exist")
	}

	return accountUser, nil
}

func (a *AcountWithDrawalService) convertObject(nickName string, object map[string]interface{}) *entities.Account {

	var account AccountData

	err := mapstructure.Decode(object[nickName], &account)
	if err != nil {
		return nil
	}

	accountDomain := entities.NewAccount(account.NickName, account.Name, account.LastName, account.Tipe, account.Amount)
	return accountDomain
}

func (a *AcountWithDrawalService) assignNewBalance(accountUser *entities.Account, newAmount float64) *entities.Account {

	accountUser.SetAmount(newAmount)

	return accountUser
}
