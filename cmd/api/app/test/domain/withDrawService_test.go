package domain

import (
	account "VERITRAN/internal/domain/account/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithDrawService(t *testing.T) {
	//Given
	depositServiceDomain := account.NewWithDrawService()
	//When
	newBalance, err := depositServiceDomain.ValidWithDraw(100, 10)
	//then
	assert.EqualValues(t, 90, newBalance)
	assert.Nil(t, err)
}

func TestWithDrawService_Negative(t *testing.T) {
	//Given
	depositServiceDomain := account.NewWithDrawService()
	//when
	_, err := depositServiceDomain.ValidWithDraw(100, -10)
	//then
	assert.Error(t, err)
}

func TestWithDrawService_funds(t *testing.T) {
	//Given
	depositServiceDomain := account.NewWithDrawService()
	//When
	_, err := depositServiceDomain.ValidWithDraw(100, 110)
	//Then
	assert.Error(t, err)
}
