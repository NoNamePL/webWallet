package models

type Wallet struct {
	valletId string `json:"valletId"`
	operationType string `json:"operationType"`
	amount float64 `json:"amount"`
}