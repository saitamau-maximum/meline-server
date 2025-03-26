package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type IMessagePresenter interface {
	GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *response.GetByChannelIDResponse
	GenerateCreateMessageResponse(message *entity.Message) *response.CreateMessageResponse
}
