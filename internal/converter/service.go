package converter

import (
	"MessageService/internal/interfaces/service"
	v1chat "MessageService/proto/message_service/chats/v1"
	v1messages "MessageService/proto/message_service/message/v1"
	v1users "MessageService/proto/message_service/users/v1"
)

// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
//
//go:generate goverter gen ./
type ServiceConverter interface {
	AddChatToService(request *v1chat.AddChatRequest) *service.AddChatRequest
	// goverter:ignore UserId
	GetChatsToService(request *v1chat.GetChatsRequest) *service.GetChatsRequest
	GetChatsToHandler(response *service.GetChatsResponse) *v1chat.GetChatsResponse

	GetMessagesToService(request *v1messages.GetMessagesRequest) *service.GetMessagesRequest
	GetMessagesToHandler(response *service.GetMessagesResponse) *v1messages.GetMessagesResponse
	SendMessageToService(request *v1messages.SendMessageRequest) *service.SendMessageRequest
	UpdateMessageStatusToService(request *v1messages.UpdateMessageStatusRequest) *service.UpdateMessageStatusRequest

	GetStatusInfoToService(request *v1users.GetStatusInfoRequest) *service.GetStatusInfoRequest
	GetStatusInfoToHandler(response *service.GetStatusInfoResponse) *v1users.GetStatusInfoResponse
	UpdateStatusToService(request *v1users.UpdateStatusRequest) *service.UpdateStatusRequest
}
