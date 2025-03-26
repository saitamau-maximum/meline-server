package request

import (
	"errors"
	"unicode/utf8"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/request"
)

type CreateMessageRequest struct {
	request.CreateMessageRequest
}

func (r *CreateMessageRequest) Validate() error {
	if r.Content == "" {
		return errors.New(config.ERR_EMPTY_MESSAGE)
	}

	if utf8.RuneCountInString(r.Content) > config.MAX_MESSAGE_LENGTH {
		return errors.New(config.ERR_TOO_LONG_MESSAGE)
	}

	return nil
}
