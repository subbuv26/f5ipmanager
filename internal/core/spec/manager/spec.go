package manager

import (
	"errors"
)

type (
	IPAddress string

	AllocateRequest struct {
		Label string
		Key   string
	}

	DeallocateRequest struct {
		Label string
		Key   string
	}

	AllocateResponse struct {
		IPAddress IPAddress
		Success   bool
	}

	DeallocateResponse struct {
		Success bool
	}
)

var (
	ErrorNotFound           = errors.New("not found")
	ErrorResourcesExhausted = errors.New("ran out of ip addresses")
)
