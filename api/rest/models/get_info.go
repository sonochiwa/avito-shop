package models

type GetInfoHandlerResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type ReceivedItem struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentItem struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type CoinHistory struct {
	Received []ReceivedItem `json:"received"`
	Sent     []SentItem     `json:"sent"`
}
