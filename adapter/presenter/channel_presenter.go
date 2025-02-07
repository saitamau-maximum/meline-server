package presenter

import (
	"strconv"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/base"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type ChannelPresenter struct{}

func NewChannelPresenter() presenter.IChannelPresenter {
	return &ChannelPresenter{}
}

func (p *ChannelPresenter) GenerateGetAllChannelsResponse(channels []*entity.Channel) *response.GetAllChannelsResponse {
	channelsResponse := &response.GetAllChannelsResponse{
		Channels: []*base.Channel{},
	}
	for _, channel := range channels {
		channelsResponse.Channels = append(channelsResponse.Channels, &base.Channel{
			Id:   strconv.FormatUint(channel.ID, 10),
			Name: channel.Name,
		})
	}

	return channelsResponse
}

func (p *ChannelPresenter) GenerateGetChannelByIdResponse(channel *entity.Channel) *response.GetChannelByIDResponse {
	childChannels := make([]*base.Channel, 0)
	for _, childChannel := range channel.ChildChannels {
		childChannels = append(childChannels, &base.Channel{
			Id:   strconv.FormatUint(childChannel.ID, 10),
			Name: childChannel.Name,
		})
	}

	users := make([]*base.User, 0)
	for _, user := range channel.Users {
		users = append(users, &base.User{
			Id:       strconv.FormatUint(user.ID, 10),
			Name:     user.Name,
			ImageUrl: user.ImageURL,
		})
	}

	return &response.GetChannelByIDResponse{
		Channel: &base.ChannelDetail{
			Id:       strconv.FormatUint(channel.ID, 10),
			Name:     channel.Name,
			Users:    users,
			Channels: childChannels,
		},
	}
}
