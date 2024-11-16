package service

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
	UserId   int64
	Username string
	Status   Status
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
	GetMessages(request *GetMessagesRequest) (*GetMessagesResponse, error)
	SendMessage(request *SendMessageRequest) error
	UpdateMessageStatus(request *UpdateMessageStatusRequest) error

	AddChat(request *AddChatRequest) error
	GetChats(request *GetChatsRequest) (*GetChatsResponse, error)

	GetStatusInfo(request *GetStatusInfoRequest) (*GetStatusInfoResponse, error)
	UpdateStatus(request *UpdateStatusRequest) error
}
