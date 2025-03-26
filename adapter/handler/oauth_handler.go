package handler

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	STATE_LENGTH = 32
)

type OAuthHandler struct {
	githubOAuthInteractor usecase.IGithubOAuthInteractor
	authInteractor        usecase.IAuthInteractor
	userInteractor        usecase.IUserInteractor
}

func NewOAuthHandler(authGroup *echo.Group, githubOAuthInteractor usecase.IGithubOAuthInteractor, authInteractor usecase.IAuthInteractor, userInteractor usecase.IUserInteractor) {
	oAuthHandler := &OAuthHandler{
		githubOAuthInteractor: githubOAuthInteractor,
		authInteractor:        authInteractor,
		userInteractor:        userInteractor,
	}

	authGroup.GET("/login", oAuthHandler.Login)
	authGroup.GET("/callback", oAuthHandler.CallBack)
	authGroup.GET("/token", oAuthHandler.GetToken)
}

func (h *OAuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	state := h.authInteractor.GenerateState(STATE_LENGTH)

	stateCookie := h.authInteractor.GenerateStateCookie(state, config.IsDev)
	c.SetCookie(stateCookie)

	url := h.githubOAuthInteractor.GetGithubOAuthURL(ctx, state)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

type OTTItem struct {
	Token string
	Exp   time.Time
}

var OTTMap = map[string]OTTItem{}
var OTTMapLock = &sync.RWMutex{}

const OAUTH_OTT_EXPIRE_CHECK_INTERVAL = 100 * time.Millisecond

func InitOTTExpireChecker() {
	ticker := time.NewTicker(OAUTH_OTT_EXPIRE_CHECK_INTERVAL)

	for range ticker.C {
		for k, v := range OTTMap {
			if time.Now().After(v.Exp) {
				delete(OTTMap, k)
			}
		}
	}
}

func (h *OAuthHandler) CallBack(c echo.Context) error {
	ctx := c.Request().Context()

	// Check State
	state := c.QueryParam("state")
	stateCookie, err := c.Cookie(config.OAUTH_STATE_COOKIE_NAME)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	if state != stateCookie.Value {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	code := c.QueryParam("code")
	gitToken, err := h.githubOAuthInteractor.GetGithubOAuthToken(ctx, code)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	userRes, err := h.githubOAuthInteractor.GetGithubUser(ctx, gitToken)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	var userId uint64

	getUserRes, err := h.userInteractor.GetUserByGithubIDOrCreate(ctx, userRes.OAuthUserID, userRes.Name, userRes.ImageURL)
	if err != nil {
		return err
	}

	userId, err = strconv.ParseUint(string(getUserRes.Id), 10, 64)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	jwt, err := h.authInteractor.GenerateJWTWithUserID(ctx, userId)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	ott := h.authInteractor.GenerateOTT(ctx)
	ottExpireNum, err := strconv.ParseInt(config.OAUTH_OTT_EXPIRE, 10, 64)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	OTTMapLock.Lock()
	OTTMap[ott] = OTTItem{
		Token: jwt,
		Exp:   time.Now().Add(time.Duration(ottExpireNum) * time.Millisecond),
	}
	OTTMapLock.Unlock()

	return c.Redirect(http.StatusTemporaryRedirect, config.FRONT_OAUTH_SUCCESS_URL+"?ott="+ott)
}

func (h *OAuthHandler) GetToken(c echo.Context) error {
	ott := c.QueryParam("ott")

	OTTMapLock.RLock()
	jwtItem, ok := OTTMap[ott]
	OTTMapLock.RUnlock()
	if !ok {
		return c.JSON(http.StatusBadRequest, "Invalid OTT")
	}

	OTTMapLock.Lock()
	delete(OTTMap, ott)
	OTTMapLock.Unlock()

	return c.JSON(http.StatusOK, map[string]string{
		"token": jwtItem.Token,
	})
}
