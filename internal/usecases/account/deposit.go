package account

import (
	"VERITRAN/internal/domain/account/entities"
	"VERITRAN/internal/domain/account/repository"
	"errors"

	serviceDomain "VERITRAN/internal/domain/account/services"

	"github.com/mitchellh/mapstructure"
)

type Account interface {
	Deposit(nickName string, amount float64) (float64, error)
}

type AcountDepositService struct {
	serviceDomain serviceDomain.DepositService
	repository    repository.AccountRepository
}

type DataMap struct {
	Data AccountData
}

type AccountData struct {
	NickName string
	Name     string
	LastName string
	Tipe     string
	Amount   float64
	Active   bool
}

func NewDepositUseCase(serviceDomain serviceDomain.DepositService, repository repository.AccountRepository) *AcountDepositService {
	return &AcountDepositService{
		serviceDomain: serviceDomain,
		repository:    repository,
	}
}

func (a *AcountDepositService) Deposit(nickName string, amount float64) (float64, error) {

	accountUser, err := a.getAccount(nickName)
	if err != nil {
		return 0, err
	}
	//Get valid bussiness domain
	newAmount, err := a.serviceDomain.ValidConsignment(*accountUser.GetAmount(), amount)
	if err != nil {
		return 0, err
	}

	accountXuser := a.assignNewBalance(accountUser, newAmount)

	result, err := a.repository.Deposit(*accountUser.GetNickName(), accountXuser)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		nickName: result,
	}

	objectFinal := a.convertObject(nickName, data)
	if objectFinal.GetAmount() == nil {
		return 0, errors.New("error with amount value")
	}

	return *objectFinal.GetAmount(), nil
}

func (a *AcountDepositService) convertObject(nickName string, object map[string]interface{}) *entities.Account {

	var account AccountData

	err := mapstructure.Decode(object[nickName], &account)
	if err != nil {
		return nil
	}

	accountDomain := entities.NewAccount(account.NickName, account.Name, account.LastName, account.Tipe, account.Amount)
	return accountDomain
}

func (a *AcountDepositService) getAccount(nickName string) (*entities.Account, error) {
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

func (a *AcountDepositService) assignNewBalance(accountUser *entities.Account, newAmount float64) *entities.Account {

	accountUser.SetAmount(newAmount)

	return accountUser
}
