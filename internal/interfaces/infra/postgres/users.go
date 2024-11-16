package postgres

import "MessageService/internal/interfaces/service"

type Users interface {
	GetStatusInfo(ids []int64) ([]*service.User, error)
	UpdateStatus(userId int64, online bool) error
}
