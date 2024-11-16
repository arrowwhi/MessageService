package postgres

import (
	"MessageService/internal/interfaces/service"
	"context"
)

type Users interface {
	GetStatusInfo(ctx context.Context, ids []int64) ([]*service.User, error)
	UpdateStatus(ctx context.Context, userId int64, online bool) error
}
