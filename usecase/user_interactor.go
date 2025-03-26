package usecase

import (
	"context"
	"database/sql"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IUserInteractor interface {
	GetUserByID(ctx context.Context, id uint64) (*response.UserMeResponse, error)
	GetUserByGithubIDOrCreate(ctx context.Context, githubID, userName, imageUrl string) (*response.UserMeResponse, error)
	IsUserExists(ctx context.Context, userID uint64) (bool, error)
}

type UserInteractor struct {
	userRepository repository.IUserRepository
	userPresenter  presenter.IUserPresenter
}

func NewUserInteractor(repository repository.IUserRepository, userPresenter presenter.IUserPresenter) IUserInteractor {
	return &UserInteractor{
		userRepository: repository,
		userPresenter:  userPresenter,
	}
}

func (i *UserInteractor) GetUserByID(ctx context.Context, id uint64) (*response.UserMeResponse, error) {
	user, err := i.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return i.userPresenter.GenerateGetUserByIdResponse(user.ToUserEntity()), nil
}

func (i *UserInteractor) GetUserByGithubIDOrCreate(ctx context.Context, githubID, userName, imageUrl string) (*response.UserMeResponse, error) {
	user, err := i.userRepository.FindByProviderID(ctx, githubID)
	if err != nil {
		if err := sql.ErrNoRows; err != nil {
			newUser := &model.User{
				ProviderID: githubID,
				Name:       userName,
				ImageURL:   imageUrl,
			}

			if _err := i.userRepository.Create(ctx, newUser); _err != nil {
				return &response.UserMeResponse{}, _err
			}

			createdUser, _err := i.userRepository.FindByProviderID(ctx, githubID)
			if _err != nil {
				return &response.UserMeResponse{}, _err
			}

			user = createdUser
		} else {
			return &response.UserMeResponse{}, err
		}
	}

	return i.userPresenter.GenerateGetUserByGithubIdResponse(user.ToUserEntity()), nil
}

func (i *UserInteractor) IsUserExists(ctx context.Context, userID uint64) (bool, error) {
	return i.userRepository.IsUserExists(ctx, userID)
}
