package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type GetUserByIdResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type GetUserByGithubIdResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type IUserPresenter interface {
	GenerateGetUserByIdResponse(user *entity.User) *response.UserMeResponse
	GenerateGetUserByGithubIdResponse(user *entity.User) *response.UserMeResponse
}
