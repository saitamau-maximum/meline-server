package request

import (
	"errors"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/request"
)

type CreateChannelRequest struct {
	request.CreateChannelRequest
}

func (r *CreateChannelRequest) Validate() error {
	if r.Name == "" {
		return errors.New(config.ERR_EMPTY_CHANNEL_NAME)
	}

	return nil
}
