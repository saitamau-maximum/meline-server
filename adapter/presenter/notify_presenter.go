package presenter

import (
	"strconv"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/base"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type NotifyPresenter struct{}

func NewNotifyPresenter() presenter.INotifyPresenter {
	return &NotifyPresenter{}
}

func (p *NotifyPresenter) GenerateNotifyMessageResponse(message *entity.Message) *response.NotifyResponse {
	return &response.NotifyResponse{
		NotifyMeta: &base.NotifyMeta{
			TypeId: config.NOTIFY_MESSAGE,
		},
		Payload: &base.Payload{
			Message: &base.Message{
				Id: message.ID,
				User: &base.User{
					Id:       strconv.FormatUint(message.User.ID, 10),
					Name:     message.User.Name,
					ImageUrl: message.User.ImageURL,
				},
				Content:        message.Content,
				ReplyToMessage: nil,
				CreatedAt:      message.CreatedAt.String(),
				UpdatedAt:      message.UpdatedAt.String(),
			},
			ChannelId: strconv.FormatUint(message.ChannelID, 10),
		},
	}
}
