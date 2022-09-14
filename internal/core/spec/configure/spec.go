package configure

import (
	"errors"
)

type (
	IPRange struct {
		Label  string
		Config string
	}

	AddIPRangeRequest struct {
		IPRange
	}

	RemoveIPRangeRequest struct {
		Label string
	}

	Response struct {
		Success bool
	}
)

var (
	ErrorInvalidInput = errors.New("invalid configuration")
	ErrorNotFound     = errors.New("not found")
)
