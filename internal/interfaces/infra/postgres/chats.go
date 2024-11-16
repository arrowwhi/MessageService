package postgres

import (
	"MessageService/internal/interfaces/service"
	"context"
)

type Chats interface {
	AddChat(ctx context.Context, chat service.Chat) error
	GetChats(ctx context.Context, userId, limit, offset int64) ([]service.Chat, error)
}
