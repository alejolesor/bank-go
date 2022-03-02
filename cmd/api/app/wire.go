package app

import (
	domain "VERITRAN/internal/domain/account/services"
	"VERITRAN/internal/platform/handler"
	"VERITRAN/internal/platform/storage/database/driver"
	repo "VERITRAN/internal/platform/storage/database/repository/account"
	"VERITRAN/internal/usecases/account"
)

func DependencyInjection() *handler.Operations {

	serviceDomain := domain.NewDepositService()
	serviceWithDrawal := domain.NewWithDrawService()

	driverSql := driver.NewDriverSql()

	repository := repo.NewAccountRepository(driverSql)

	useCaseDeposit := account.NewDepositUseCase(*serviceDomain, repository)
	useCaseCreate := account.NewCreateUseCase(repository)
	useCaseWithDrwal := account.NewWithDrawUseCase(*serviceWithDrawal, repository)
	useCaseTransfer := account.NewTransferUseCase(*serviceWithDrawal, *serviceDomain, repository)

	Operations := handler.NewOperations(useCaseDeposit, useCaseCreate, useCaseWithDrwal, useCaseTransfer)

	return Operations
}
