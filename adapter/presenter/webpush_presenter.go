package presenter

import (
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type WebPushPresneter struct{}

func NewWebPushPresenter() presenter.IWebPushPresenter {
	return &WebPushPresneter{}
}

func (p *WebPushPresneter) GetPublicKeyResponse(pubKey string) *response.GetPublicKeyResponse {
	return &response.GetPublicKeyResponse{
		PublicKey: pubKey,
	}
}
