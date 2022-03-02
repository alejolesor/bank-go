package usecases

import (
	mockDomain "VERITRAN/cmd/api/app/test/domain/mocks"
	"VERITRAN/internal/domain/account/entities"
	serviceDomain "VERITRAN/internal/domain/account/services"
	"VERITRAN/internal/usecases/account"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepository mockDomain.AccountRepository

func TestTransfer(t *testing.T) {
	transfer := entities.Transfer{Depositor: *entities.NewDepositor("francisco", 2000), Receiver: *entities.NewReceiver("alejo")}

	structAccount := account.AccountData{
		NickName: "francisco",
		Amount:   4000,
	}
	structAccount2 := account.AccountData{
		NickName: "alejo",
		Amount:   2000,
	}
	var accountMap = map[string]interface{}{
		"francisco": structAccount,
	}

	var accountMap2 = map[string]interface{}{
		"alejo": structAccount2,
	}

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			//Given
			name: "Success Method Deposit",
			mock: func() {
				mockRepository.On("Get", "francisco").Return(&accountMap, nil)
				mockRepository.On("Get", "alejo").Return(&accountMap2, nil)
				mockRepository.On("Deposit", "alejo", mock.Anything).Return(mock.Anything, nil)
				mockRepository.On("WithDrawal", "francisco", mock.Anything).Return(float32(2000), nil)

			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepository = mockDomain.AccountRepository{}

			tt.mock()
			servicesDomain := serviceDomain.NewDepositService()
			serviceWithDrawal := serviceDomain.NewWithDrawService()
			useCase := account.NewTransferUseCase(*serviceWithDrawal, *servicesDomain, &mockRepository)
			//when
			_, err := useCase.Transfer(transfer)
			//then

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
