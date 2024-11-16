package postgres

import "MessageService/internal/interfaces/service"

type Chats interface {
	AddChat(chat service.Chat) error
	GetChats(userId, limit, offset int64) ([]service.Chat, error)
}
