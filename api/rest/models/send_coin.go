package models

type SendCoinHandlerRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
