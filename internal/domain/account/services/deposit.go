package account

import (
	"errors"
)

type DepositService struct {
}

func NewDepositService() *DepositService {
	return &DepositService{}
}

func (d *DepositService) ValidConsignment(valueCurrently, deposit float64) (float64, error) {
	if deposit < 0 {
		return 0, errors.New("negative value is invalid")
	}
	valueCurrently += deposit
	return valueCurrently, nil
}
