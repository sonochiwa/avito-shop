package converter

import (
	"github.com/sonochiwa/avito-shop/api/rest/models"
	models2 "github.com/sonochiwa/avito-shop/internal/domain"
)

func ConvertGetInfoToResponse(info models2.GetInfo) models.GetInfoHandlerResponse {
	inventory := make([]models.InventoryItem, len(info.Inventory))
	for i, item := range info.Inventory {
		inventory[i] = models.InventoryItem{
			Type:     item.ProductTitle,
			Quantity: item.Quantity,
		}
	}

	coinHistory := models.CoinHistory{
		Received: make([]models.ReceivedItem, len(info.Ledger.Received)),
		Sent:     make([]models.SentItem, len(info.Ledger.Sent)),
	}

	for i, received := range info.Ledger.Received {
		coinHistory.Received[i] = models.ReceivedItem{
			FromUser: received.FromUser,
			Amount:   received.Amount,
		}
	}

	for i, sent := range info.Ledger.Sent {
		coinHistory.Sent[i] = models.SentItem{
			ToUser: sent.ToUser,
			Amount: sent.Amount,
		}
	}

	return models.GetInfoHandlerResponse{
		Coins:       info.Balance,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}
}
