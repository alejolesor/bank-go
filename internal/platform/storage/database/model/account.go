package model

import "github.com/google/uuid"

type Account struct {
	Id       int
	NickName string
	Name     string
	LastName string
	Tipe     string
	Amount   float32
	Active   bool
}

func NewAccountModel(nickName, name, lastName, tipe string, amount float32, active bool) *Account {
	return &Account{
		Id:       int(uuid.New().ID()),
		NickName: nickName,
		Name:     name,
		LastName: lastName,
		Tipe:     tipe,
		Amount:   amount,
		Active:   active,
	}
}
