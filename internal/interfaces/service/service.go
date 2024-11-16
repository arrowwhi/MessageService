package service

import "context"

type Image struct {
	ImageId int64
	Data    []byte
}

type Status struct {
	IsOnline       bool
	LastActiveTime int64
}

type Message struct {
	MessageId   int64
	SenderId    int64
	RecipientId int64
	Text        string
	Images      []Image
	Timestamp   int64
}

type Chat struct {
	Id    int64
	Name  string
	Users []int64
}

type User struct {
	UserId     int64
	Username   string
	Status     Status
	LastActive int64
}

type GetMessagesRequest struct {
	ChatId int64
	Limit  int64
	Offset int64
}

type GetMessagesResponse struct {
	Messages []Message
}

type SendMessageRequest struct {
	Message Message
}

type UpdateMessageStatusRequest struct {
	MessageId int64
}

type AddChatRequest struct {
	Chat Chat
}

type GetChatsRequest struct {
	UserId int64
	Limit  int64
	Offset int64
}

type GetChatsResponse struct {
	Chats []Chat
}

type GetStatusInfoRequest struct {
	Ids []int64
}

type GetStatusInfoResponse struct {
	Users []*User
}

type UpdateStatusRequest struct {
	UserId int64
	Online bool
}

type Service interface {
	GetMessages(ctx context.Context, request *GetMessagesRequest) (*GetMessagesResponse, error)
	SendMessage(ctx context.Context, request *SendMessageRequest) error
	UpdateMessageStatus(ctx context.Context, request *UpdateMessageStatusRequest) error

	AddChat(ctx context.Context, request *AddChatRequest) error
	GetChats(ctx context.Context, request *GetChatsRequest) (*GetChatsResponse, error)

	GetStatusInfo(ctx context.Context, request *GetStatusInfoRequest) (*GetStatusInfoResponse, error)
	UpdateStatus(ctx context.Context, request *UpdateStatusRequest) error
}
