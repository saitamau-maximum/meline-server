package presenter

import (
	"strconv"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type MessagePresenter struct{}

func NewMessagePresenter() presenter.IMessagePresenter {
	return &MessagePresenter{}
}

func (p *MessagePresenter) GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *presenter.GetMessagesByChannelIDResponse {
	messagesResponse := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{},
	}
	for _, message := range messages {
		var replyToMessage *presenter.ReplyToMessage = nil
		if message.ReplyToMessage != nil {
			replyToMessage = &presenter.ReplyToMessage{
				ID: message.ReplyToMessage.ID,
				User: &presenter.User{
					ID:       strconv.FormatUint(message.ReplyToMessage.User.ID, 10),
					Name:     message.ReplyToMessage.User.Name,
					ImageURL: message.ReplyToMessage.User.ImageURL,
				},
				Content: message.ReplyToMessage.Content,
			}
		}
		messagesResponse.Messages = append(messagesResponse.Messages, &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       strconv.FormatUint(message.User.ID, 10),
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}

func (p *MessagePresenter) GenerateCreateMessageResponse(message *entity.Message) *presenter.CreateMessageResponse {
	var replyToMessage *presenter.ReplyToMessage = nil
	if message.ReplyToMessage != nil {
		replyToMessage = &presenter.ReplyToMessage{
			ID: message.ReplyToMessage.ID,
			User: &presenter.User{
				ID:       strconv.FormatUint(message.ReplyToMessage.User.ID, 10),
				Name:     message.ReplyToMessage.User.Name,
				ImageURL: message.ReplyToMessage.User.ImageURL,
			},
			Content: message.ReplyToMessage.Content,
		}
	}

	return &presenter.CreateMessageResponse{
		Message: &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       strconv.FormatUint(message.User.ID, 10),
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		},
		ChannelID: strconv.FormatUint(message.ChannelID, 10),
	}
}
