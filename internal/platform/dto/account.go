package dto

type Account struct {
	NickName string  `json:"nickName"`
	Name     string  `json:"name"`
	LastName string  `json:"lastName"`
	Tipe     string  `json:"type"`
	Amount   float64 `json:"amount"`
	Active   bool    `json:"active"`
}

type Deposit struct {
	NickName string  `json:"nickName"`
	Amount   float64 `json:"amount"`
}

type WithDrawal struct {
	NickName   string  `json:"nickName"`
	WithDrawal float64 `json:"withdrawal"`
}

type Transfer struct {
	Depositor Depositor `json:"depositor"`
	Receiver  Receiver  `json:"receiver"`
}

type Depositor struct {
	NickName string  `json:"nickName"`
	Amount   float64 `json:"amount"`
}

type Receiver struct {
	NickName string `json:"nickName"`
}
