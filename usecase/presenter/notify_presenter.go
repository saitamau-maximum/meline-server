package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type INotifyPresenter interface {
	GenerateNotifyMessageResponse(message *entity.Message) *response.NotifyResponse
}
