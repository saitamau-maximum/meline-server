package handler

import (
	"net/http"
	"strconv"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/request"
	"github.com/saitamau-maximum/meline/usecase"
)

type WebPushHandler struct {
	pushServiceInteractor usecase.IWebPushInteractor
}

func NewWebPushHandler(webPushGroup *echo.Group, pushServiceInteractor usecase.IWebPushInteractor) {
	webPushHandler := &WebPushHandler{
		pushServiceInteractor: pushServiceInteractor,
	}

	webPushGroup.POST("/subscription", webPushHandler.StoreSubscription)
	webPushGroup.GET("/publickey", webPushHandler.GetPublicKey)
	webPushGroup.DELETE("/subscription/:key", webPushHandler.DeleteSubscription)
}

func (h *WebPushHandler) GetPublicKey(c echo.Context) error {
	pubKey := h.pushServiceInteractor.GetPublicKey()
	return c.JSON(http.StatusOK, pubKey)
}

func (h *WebPushHandler) StoreSubscription(c echo.Context) error {
	req := new(request.StoreSubscriptionRequest)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	subscription := &webpush.Subscription{
		Endpoint: req.Endpoint,
		Keys: webpush.Keys{
			P256dh: req.P256Dh,
			Auth:   req.Auth,
		},
	}

	userId := c.Get("user_id").(uint64)

	if err := h.pushServiceInteractor.StoreSubscription(c.Request().Context(), strconv.FormatUint(userId, 10), subscription); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *WebPushHandler) DeleteSubscription(c echo.Context) error {
	key := c.Param("key")

	if err := h.pushServiceInteractor.DeleteSubscription(c.Request().Context(), key); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
