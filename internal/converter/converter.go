package converter

import (
	"github.com/sonochiwa/avito-shop/api/rest"
	"github.com/sonochiwa/avito-shop/internal/domain"
)

func ConvertGetInfoToResponse(info domain.GetInfo) rest.GetInfoHandlerResponse {
	inventory := make([]rest.InventoryItem, len(info.Inventory))
	for i, item := range info.Inventory {
		inventory[i] = rest.InventoryItem{
			Type:     item.ProductTitle,
			Quantity: item.Quantity,
		}
	}

	coinHistory := rest.CoinHistory{
		Received: make([]rest.ReceivedItem, len(info.Ledger.Received)),
		Sent:     make([]rest.SentItem, len(info.Ledger.Sent)),
	}

	for i, received := range info.Ledger.Received {
		coinHistory.Received[i] = rest.ReceivedItem{
			FromUser: received.FromUser,
			Amount:   received.Amount,
		}
	}

	for i, sent := range info.Ledger.Sent {
		coinHistory.Sent[i] = rest.SentItem{
			ToUser: sent.ToUser,
			Amount: sent.Amount,
		}
	}

	return rest.GetInfoHandlerResponse{
		Coins:       info.Balance,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}
}
