package presenter

import (
	"strconv"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type UserPresenter struct{}

func NewUserPresenter() presenter.IUserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) GenerateGetUserByIdResponse(user *entity.User) *response.UserMeResponse {
	return &response.UserMeResponse{
		Id:       strconv.FormatUint(user.ID, 10),
		Name:     user.Name,
		ImageUrl: user.ImageURL,
	}
}

func (p *UserPresenter) GenerateGetUserByGithubIdResponse(user *entity.User) *response.UserMeResponse {
	return &response.UserMeResponse{
		Id:       strconv.FormatUint(user.ID, 10),
		Name:     user.Name,
		ImageUrl: user.ImageURL,
	}
}
