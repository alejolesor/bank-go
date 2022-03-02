package entities

type Transfer struct {
	Depositor Depositor
	Receiver  Receiver
}

type Depositor struct {
	nickName string
	amount   float64
	balance  float64
}

type Receiver struct {
	nickName string
	balance  float64
}

func NewTransfer(depositorNickName string, amountTransference float64, receiverNickName string) *Transfer {
	return &Transfer{
		Depositor: Depositor{nickName: depositorNickName, amount: amountTransference},
		Receiver:  Receiver{nickName: receiverNickName},
	}
}
func NewDepositor(nickName string, amount float64) *Depositor {
	return &Depositor{nickName: nickName, amount: amount}
}
func NewReceiver(nickName string) *Receiver {
	return &Receiver{nickName: nickName}
}

func (a *Depositor) GetNickNameDepositor() *string {
	return &a.nickName
}

func (a *Depositor) GetAmuntDeposit() *float64 {
	return &a.amount
}

func (r *Receiver) GetNickNameReceiver() *string {
	return &r.nickName
}
