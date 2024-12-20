package gateway

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/adapter/response"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase"
)

type AuthGateway struct {
	userInteractor usecase.IUserInteractor
}

func NewAuthGateway(userInteractor usecase.IUserInteractor) *AuthGateway {
	return &AuthGateway{
		userInteractor: userInteractor,
	}
}

func (h *AuthGateway) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		}

		accessToken := authHeader[len("Bearer "):]
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unsupported signing method")
			}

			return []byte(config.JWT_SECRET), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		}

		userId := claims["user_id"].(float64)

		isExistUser, err := h.userInteractor.IsUserExists(c.Request().Context(), uint64(userId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "Internal Server Error"})
		} else if !isExistUser {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		}

		c.Set("user_id", uint64(userId))

		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		}

		return next(c)
	}
}
