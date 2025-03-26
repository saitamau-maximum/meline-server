package presenter

import (
	"strconv"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/base"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type MessagePresenter struct{}

func NewMessagePresenter() presenter.IMessagePresenter {
	return &MessagePresenter{}
}

func (p *MessagePresenter) GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *response.GetByChannelIDResponse {
	messagesResponse := &response.GetByChannelIDResponse{
		Messages: []*base.Message{},
	}
	for _, message := range messages {
		var replyToMessage *base.ReplyToMessage = nil
		if message.ReplyToMessage != nil {
			replyToMessage = &base.ReplyToMessage{
				Id: message.ReplyToMessage.ID,
				User: &base.User{
					Id:       strconv.FormatUint(message.ReplyToMessage.User.ID, 10),
					Name:     message.ReplyToMessage.User.Name,
					ImageUrl: message.ReplyToMessage.User.ImageURL,
				},
				Content: message.ReplyToMessage.Content,
			}
		}
		messagesResponse.Messages = append(messagesResponse.Messages, &base.Message{
			Id: message.ID,
			User: &base.User{
				Id:       strconv.FormatUint(message.User.ID, 10),
				Name:     message.User.Name,
				ImageUrl: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}

func (p *MessagePresenter) GenerateCreateMessageResponse(message *entity.Message) *response.CreateMessageResponse {
	return &response.CreateMessageResponse{}
}
