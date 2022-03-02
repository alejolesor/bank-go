package account

import "errors"

type WithDrawService struct {
}

func NewWithDrawService() *WithDrawService {
	return &WithDrawService{}
}

func (w *WithDrawService) ValidWithDraw(valueCurrently, withDraw float64) (float64, error) {
	if withDraw < 0 {
		return 0, errors.New("negative value is invalid")
	}
	if withDraw > valueCurrently {
		return 0, errors.New("insufficient funds")
	}
	valueCurrently -= withDraw
	return valueCurrently, nil
}
