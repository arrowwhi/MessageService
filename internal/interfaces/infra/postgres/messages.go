package postgres

import "MessageService/internal/interfaces/service"

type Message interface {
	GetMessages(chatId, limit, offset int64) ([]service.Message, error)
	AddMessage(message service.Message) error
	UpdateMessageStatus(messageId int64, read bool) error
}
