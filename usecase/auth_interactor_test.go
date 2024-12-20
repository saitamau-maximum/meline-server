package usecase_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/saitamau-maximum/meline/usecase"

	"github.com/stretchr/testify/assert"
)

func TestAuthInteractor_GenerateJWT(t *testing.T) {
	interactor := usecase.NewAuthInteractor()

	var userId uint64 = 1

	result, err := interactor.GenerateJWTWithUserID(context.Background(), userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestAuthInteractor_GenerateState(t *testing.T) {
	interactor := usecase.NewAuthInteractor()

	result := interactor.GenerateState(10)
	assert.Len(t, result, 10*2)
}

func TestAuthInteractor_GenerateStateCookie(t *testing.T) {
	interactor := usecase.NewAuthInteractor()

	cookie := interactor.GenerateStateCookie("this_is_state", false)
	assert.Equal(t, "state", cookie.Name)
	assert.Equal(t, "this_is_state", cookie.Value)
	assert.Equal(t, "/", cookie.Path)
	assert.True(t, cookie.HttpOnly)
	assert.Equal(t, http.SameSiteLaxMode, cookie.SameSite)
	assert.True(t, cookie.Secure)
	assert.True(t, cookie.Expires.After(time.Now()))
}
