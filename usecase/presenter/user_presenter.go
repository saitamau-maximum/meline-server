package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
)

type IUserPresenter interface {
	GenerateGetUserByIdResponse(user *entity.User) *response.UserMeResponse
	GenerateGetUserByGithubIdResponse(user *entity.User) *response.UserMeResponse
}
