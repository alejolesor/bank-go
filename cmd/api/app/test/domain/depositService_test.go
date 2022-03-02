package domain

import (
	account "VERITRAN/internal/domain/account/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepositService(t *testing.T) {
	//Given
	depositServiceDomain := account.NewDepositService()
	//when
	newBalance, err := depositServiceDomain.ValidConsignment(100, 10)
	//then
	assert.EqualValues(t, 110, newBalance)
	assert.Nil(t, err)
}

func TestDepositNegative(t *testing.T) {
	//Given
	depositServiceDomain := account.NewDepositService()
	//When
	_, err := depositServiceDomain.ValidConsignment(100, -10)
	//Then
	assert.Error(t, err)
}
