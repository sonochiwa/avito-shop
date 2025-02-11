package converter

import (
	"github.com/sonochiwa/avito-shop/api/rest"
	"github.com/sonochiwa/avito-shop/internal/domain"
)

func ConvertGetInfoToResponse(info domain.GetInfo) rest.GetInfoHandlerResponse {
	inventory := make([]rest.InventoryItem, 0, len(info.Inventory))
	for _, item := range info.Inventory {
		inventory = append(inventory, rest.InventoryItem{
			Type:     item.ProductTitle,
			Quantity: item.Quantity,
		})
	}

	coinHistory := rest.CoinHistory{
		Received: make([]rest.ReceivedItem, 0, len(info.Ledger.Received)),
		Sent:     make([]rest.SentItem, 0, len(info.Ledger.Sent)),
	}

	for _, received := range info.Ledger.Received {
		coinHistory.Received = append(coinHistory.Received, rest.ReceivedItem{
			FromUser: received.FromUser,
			Amount:   received.Amount,
		})
	}

	for _, sent := range info.Ledger.Sent {
		coinHistory.Sent = append(coinHistory.Sent, rest.SentItem{
			ToUser: sent.ToUser,
			Amount: sent.Amount,
		})
	}

	return rest.GetInfoHandlerResponse{
		Coins:       info.Balance,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}
}
