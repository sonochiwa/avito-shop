package domain

type GetInfo struct {
	Balance   int             `json:"balance"`
	Inventory []InventoryItem `json:"inventory"`
	Ledger    Ledger          `json:"ledger"`
}

type InventoryItem struct {
	ProductTitle string `json:"product_title"`
	Quantity     int    `json:"quantity"`
}

type ReceivedItem struct {
	FromUser string `json:"from_user"`
	Amount   int    `json:"amount"`
}

type SentItem struct {
	ToUser string `json:"to_user"`
	Amount int    `json:"amount"`
}

type Ledger struct {
	Received []ReceivedItem `json:"received"`
	Sent     []SentItem     `json:"sent"`
}
