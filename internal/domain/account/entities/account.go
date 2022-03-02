package entities

type Account struct {
	nickname string
	name     string
	lastname string
	tipe     string
	amount   float64
	active   bool
}

func NewAccount(nickname, name, lastname, tipe string, amount float64) *Account {
	return &Account{
		nickname: nickname,
		name:     name,
		lastname: lastname,
		tipe:     tipe,
		amount:   amount,
		active:   true,
	}
}

func (a *Account) GetNickName() *string {
	return &a.nickname
}

func (a *Account) Getname() *string {
	return &a.name
}

func (a *Account) GetLastName() *string {
	return &a.lastname
}

func (a *Account) GetType() *string {
	return &a.tipe
}

func (a *Account) GetAmount() *float64 {
	return &a.amount
}

func (a *Account) GetActive() *bool {
	return &a.active
}

func (a *Account) SetAmount(newamount float64) bool {
	a.amount = newamount
	return true
}
