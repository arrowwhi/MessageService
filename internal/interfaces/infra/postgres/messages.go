package postgres

import (
	"MessageService/internal/interfaces/service"
	"context"
)

type Message interface {
	GetMessages(ctx context.Context, chatId, limit, offset int64) ([]service.Message, error)
	AddMessage(ctx context.Context, message service.Message) error
	UpdateMessageStatus(ctx context.Context, messageId int64, read bool) error
}
