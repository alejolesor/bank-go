package usecases

import (
	"testing"

	mockDomain "VERITRAN/cmd/api/app/test/domain/mocks"
	"VERITRAN/cmd/api/app/test/useCases/mocks"
	serviceDomain "VERITRAN/internal/domain/account/services"
	"VERITRAN/internal/usecases/account"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockServiceAccount mocks.Account
var mockRepositoryAccount mockDomain.AccountRepository

func TestDeposit_ok(t *testing.T) {

	structAccount := account.AccountData{
		NickName: "francisco",
		Amount:   110,
	}
	var accountMap = map[string]interface{}{
		"francisco": structAccount,
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
				mockRepositoryAccount.On("Get", "francisco").Return(&accountMap, nil)
				mockRepositoryAccount.On("Deposit", "francisco", mock.Anything).Return(structAccount, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServiceAccount = mocks.Account{}
			mockRepositoryAccount = mockDomain.AccountRepository{}
			expectedBalance := float64(110)

			tt.mock()
			servicesDomain := serviceDomain.NewDepositService()
			useCase := account.NewDepositUseCase(*servicesDomain, &mockRepositoryAccount)
			//when
			newBalance, err := useCase.Deposit("francisco", 10)
			//then
			if newBalance > 0 {
				assert.EqualValues(t, expectedBalance, newBalance)
			}

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
