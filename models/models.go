package models

type Wallet struct {
	ValletId      string `json:"valletId"`
	OperationType string `json:"operationType"`
	Amount        int    `json:"amount"`
}
