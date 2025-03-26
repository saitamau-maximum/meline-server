package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type NotifyMeta struct {
	TypeID uint64 `json:"type_id"`
}

type Payload struct {
	Message   *Message `json:"message"`
	ChannelID string   `json:"channel_id"`
}

type NotifyMessageResponse struct {
	NotifyMeta NotifyMeta `json:"notify_meta"`
	Payload    Payload    `json:"payload"`
}

type INotifyPresenter interface {
	GenerateNotifyMessageResponse(message *entity.Message) *response.NotifyResponse
}
