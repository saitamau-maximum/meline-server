package usecase_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/base"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"

	"github.com/stretchr/testify/assert"
)

type FindByIDFailed string
type FindByNameFailed string
type JoinFailed string
type LeaveFailed string
type UpdateFailed string
type DeleteFailed string

const (
	FindByIDFailedValue   FindByIDFailed   = "find_by_id_failed"
	FindByNameFailedValue FindByNameFailed = "find_by_name_failed"
	JoinFailedValue       JoinFailed       = "join_failed"
	LeaveFailedValue      LeaveFailed      = "leave_failed"
	UpdateFailedValue     UpdateFailed     = "update_failed"
	DeleteFailedValue     DeleteFailed     = "delete_failed"
)

func TestChannelInteractor_Success_GetAllChannels(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	expectedChannels := []*base.Channel{
		{
			Id:   "1",
			Name: "test-channel",
		},
	}

	result, err := interactor.GetAllChannels(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannels, result.Channels)
}

func TestChannelInteractor_Failed_GetAllChannels(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, FindChannelsFailedValue, true)

	expectedChannels := &response.GetAllChannelsResponse{
		Channels: []*base.Channel{},
	}

	result, err := interactor.GetAllChannels(ctx, 1)
	assert.Error(t, err)
	assert.Equal(t, expectedChannels, result)
}

func TestChannelInteractor_Success_GetChannelByID(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	expectedChannel := &base.ChannelDetail{
		Id:  "1",
		Name: "test-channel",
		Users: []*base.User{
			{
				Id:       "1",
				Name:     "John Doe",
				ImageUrl: "https://example.com/image.jpg",
			},
		},
	}

	result, err := interactor.GetChannelByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannel, result.Channel)
}

func TestChannelInteractor_Failed_GetChannelByID(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, FindChannelsFailedValue, true)

	result, err := interactor.GetChannelByID(ctx, 2)

	expectedChannel := &response.GetChannelByIDResponse{}

	assert.Error(t, err)
	assert.Equal(t, expectedChannel, result)
}

func TestChannelInteractor_Success_CreateChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.CreateChannel(ctx, "test-channel", 1)

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_CreateChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, CreateFailedValue, true)

	err := interactor.CreateChannel(ctx, "test-channel", 1)

	assert.Error(t, err)
}

func TestChannelInteractor_Success_UpdateChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.UpdateChannel(ctx, 1, "test-channel")

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_UpdateChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, UpdateFailedValue, true)

	err := interactor.UpdateChannel(ctx, 1, "test-channel")

	assert.Error(t, err)
}

func TestChannelInteractor_Success_DeleteChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.DeleteChannel(ctx, 1)

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_DeleteChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, DeleteFailedValue, true)

	err := interactor.DeleteChannel(ctx, 1)

	assert.Error(t, err)
}

func TestChannelInteractor_Success_JoinChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.JoinChannel(ctx, 1, 1)

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_JoinChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, JoinFailedValue, true)

	err := interactor.JoinChannel(ctx, 1, 1)

	assert.Error(t, err)
}

func TestChannelInteractor_Success_LeaveChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.LeaveChannel(ctx, 1, 1)

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_LeaveChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, LeaveFailedValue, true)

	err := interactor.LeaveChannel(ctx, 1, 1)

	assert.Error(t, err)
}

func TestChannelInteractor_Success_CreateChildChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)

	err := interactor.CreateChildChannel(ctx, "test-channel", 1, 1)

	assert.NoError(t, err)
}

func TestChannelInteractor_Failed_CreateChildChannel(t *testing.T) {
	ctx := context.Background()
	hub := entity.NewHubEntity()
	repo := &mockChannelRepository{}
	repoUser := &mockUserRepository{}
	pre := &mockChannelPresenter{}

	interactor := usecase.NewChannelInteractor(hub, repo, repoUser, pre)
	ctx = context.WithValue(ctx, CreateFailedValue, true)

	err := interactor.CreateChildChannel(ctx, "test-channel", 1, 1)

	assert.Error(t, err)
}

type mockChannelRepository struct{}

func (m *mockChannelRepository) FindByID(ctx context.Context, id uint64) (*model.Channel, error) {
	if ctx.Value(FindChannelsFailedValue) != nil {
		return nil, fmt.Errorf("find all failed")
	}

	return &model.Channel{
		ID:   1,
		Name: "test-channel",
		Users: []*model.User{
			{
				ID:       1,
				Name:     "John Doe",
				ImageURL: "https://example.com/image.jpg",
			},
		},
	}, nil
}

func (m *mockChannelRepository) FindByName(ctx context.Context, name string) ([]*model.Channel, error) {
	if ctx.Value(FindByNameFailedValue) != nil {
		return nil, fmt.Errorf("find all failed")
	}

	return []*model.Channel{
		{
			ID:   1,
			Name: "test-channel",
			Users: []*model.User{
				{
					ID:       1,
					Name:     "John Doe",
					ImageURL: "https://example.com/image.jpg",
				},
			},
		},
	}, nil
}

func (m *mockChannelRepository) Create(ctx context.Context, channel *model.Channel, userID uint64) error {
	if ctx.Value(CreateFailedValue) != nil {
		return fmt.Errorf("create failed")
	}

	return nil
}

func (m *mockChannelRepository) CreateChildChannel(ctx context.Context, channel *model.Channel, parentChannelID uint64, userID uint64) error {
	if ctx.Value(CreateFailedValue) != nil {
		return fmt.Errorf("create failed")
	}

	return nil
}

func (m *mockChannelRepository) Update(ctx context.Context, channel *model.Channel) error {
	if ctx.Value(UpdateFailedValue) != nil {
		return fmt.Errorf("update failed")
	}

	return nil
}

func (m *mockChannelRepository) Delete(ctx context.Context, id uint64) error {
	if ctx.Value(DeleteFailedValue) != nil {
		return fmt.Errorf("delete failed")
	}

	return nil
}

func (m *mockChannelRepository) JoinChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if ctx.Value(JoinFailedValue) != nil {
		return fmt.Errorf("join failed")
	}

	return nil
}

func (m *mockChannelRepository) LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if ctx.Value(LeaveFailedValue) != nil {
		return fmt.Errorf("leave failed")
	}

	return nil
}

func (m *mockChannelRepository) FetchJoinedChannelIDs(ctx context.Context, userID uint64) ([]uint64, error) {
	return []uint64{1}, nil
}

type mockChannelPresenter struct{}

func (m *mockChannelPresenter) GenerateGetChannelByIdResponse(channel *entity.Channel) *response.GetChannelByIDResponse {
	users := make([]*base.User, len(channel.Users))
	for i, u := range channel.Users {
		users[i] = &base.User{
			Id:       strconv.FormatUint(u.ID, 10),
			Name:     u.Name,
			ImageUrl: u.ImageURL,
		}
	}

	return &response.GetChannelByIDResponse{
		Channel: &base.ChannelDetail{
			Id:       strconv.FormatUint(channel.ID, 10),
			Name:     channel.Name,
			Users:    users,
		},
	}
}

func (m *mockChannelPresenter) GenerateGetAllChannelsResponse(channels []*entity.Channel) *response.GetAllChannelsResponse {
	res := make([]*base.Channel, len(channels))
	for i, c := range channels {
		res[i] = &base.Channel{
			Id:   strconv.FormatUint(c.ID, 10),
			Name: c.Name,
		}
	}

	return &response.GetAllChannelsResponse{
		Channels: res,
	}
}
