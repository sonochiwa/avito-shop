package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonochiwa/avito-shop/internal/domain"
)

type Repository struct {
	conn *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Repository {
	return Repository{
		conn: pool,
	}
}

func (r Repository) GetInfo(ctx context.Context, userID int) (domain.GetInfo, error) {
	var bytes []byte
	var result domain.GetInfo

	err := r.conn.QueryRow(ctx, getInfo, userID).Scan(&bytes)
	if err != nil {
		return domain.GetInfo{}, fmt.Errorf("conn.QueryRow: %w", err)
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return domain.GetInfo{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (r Repository) SendCoin() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) BuyItem() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Auth() {
	//TODO implement me
	panic("implement me")
}
