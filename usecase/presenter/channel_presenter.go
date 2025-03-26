package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type IChannelPresenter interface {
	GenerateGetAllChannelsResponse(channels []*entity.Channel) *response.GetAllChannelsResponse
	GenerateGetChannelByIdResponse(channel *entity.Channel) *response.GetChannelByIDResponse
}
