package manager

import (
	"context"
	spec "f5ipmanager/internal/core/spec/manager"
)

// ManagerService declares the core functionalities
type ManagerService interface {
	AllocateIPAddress(ctx context.Context, request spec.AllocateRequest) (spec.AllocateResponse, error)
	DeallocateIPAddress(ctx context.Context, request spec.DeallocateRequest) (spec.DeallocateResponse, error)
}

type ManagerRepository interface {
	GetIPAddress(ctx context.Context, label, key string) (spec.IPAddress, error)
	AllocateIPAddress(ctx context.Context, label, key string) (spec.IPAddress, error)
	FreeIPAddress(ctx context.Context, label, key string) error
}
