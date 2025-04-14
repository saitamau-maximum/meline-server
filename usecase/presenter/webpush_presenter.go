package presenter

import "github.com/saitamau-maximum/meline/generated/proto/go/schema/response"

type IWebPushPresenter interface {
	GetPublicKeyResponse(pubKey string) *response.GetPublicKeyResponse
}
