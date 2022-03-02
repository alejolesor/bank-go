package account

import (
	"VERITRAN/internal/domain/account/entities"

	"VERITRAN/internal/domain/account/repository"
	serviceDomain "VERITRAN/internal/domain/account/services"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type AccountTransfer interface {
	Transfer(transfer entities.Transfer) (*string, error)
}

type AcountTransferService struct {
	serviceDomainWithDrawal serviceDomain.WithDrawService
	serviceDomainDeposit    serviceDomain.DepositService
	repository              repository.AccountRepository
}

func NewTransferUseCase(serviceDomainWithDrawal serviceDomain.WithDrawService, serviceDDomainDeposit serviceDomain.DepositService, repository repository.AccountRepository) *AcountTransferService {
	return &AcountTransferService{
		serviceDomainWithDrawal: serviceDomainWithDrawal,
		serviceDomainDeposit:    serviceDDomainDeposit,
		repository:              repository,
	}
}

func (a *AcountTransferService) Transfer(transfer entities.Transfer) (*string, error) {

	result := "process success transfer"
	//get account depositor and receiver to validat exists

	depositor, err := a.getAccount(*transfer.Depositor.GetNickNameDepositor())
	if err != nil {
		return nil, err
	}

	receiver, err := a.getAccount(*transfer.Receiver.GetNickNameReceiver())
	if err != nil {
		return nil, err
	}

	//valid balance higher than deposit  and do withdrawal at account  depositor
	newBalanceDepositor, err := a.serviceDomainWithDrawal.ValidWithDraw(*depositor.GetAmount(), *transfer.Depositor.GetAmuntDeposit())
	if err != nil {
		return nil, err
	}
	//do deposit account receiver
	newBalanceReceiver, err := a.serviceDomainDeposit.ValidConsignment(*receiver.GetAmount(), *transfer.Depositor.GetAmuntDeposit())
	if err != nil {
		return nil, err
	}
	receiver = a.assignNewBalance(receiver, newBalanceReceiver)
	depositor = a.assignNewBalance(depositor, newBalanceDepositor)

	//update account receiver
	resulDeposit, err := a.repository.Deposit(*receiver.GetNickName(), receiver)
	if err != nil {
		return nil, err
	}

	//update account depositor
	resultWithDrawal, err := a.repository.WithDrawal(*depositor.GetNickName(), depositor)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Balance Depositor  %v , Balance Receiver %v", resulDeposit, resultWithDrawal)

	return &result, nil

}

func (a *AcountTransferService) getAccount(nickName string) (*entities.Account, error) {
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

func (a *AcountTransferService) convertObject(nickName string, object map[string]interface{}) *entities.Account {

	var account AccountData

	err := mapstructure.Decode(object[nickName], &account)
	if err != nil {
		return nil
	}

	accountDomain := entities.NewAccount(account.NickName, account.Name, account.LastName, account.Tipe, account.Amount)
	return accountDomain
}

func (a *AcountTransferService) assignNewBalance(accountUser *entities.Account, newAmount float64) *entities.Account {

	accountUser.SetAmount(newAmount)

	return accountUser
}
