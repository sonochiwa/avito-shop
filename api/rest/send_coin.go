package rest

type SendCoinHandlerRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
