package configure

import (
	"context"
	spec "f5ipmanager/internal/core/spec/configure"
)

// ConfigService declares the core functionalities
type ConfigService interface {
	AddIPRange(ctx context.Context, request spec.AddIPRangeRequest) (spec.Response, error)
	RemoveIPRange(ctx context.Context, request spec.RemoveIPRangeRequest) (spec.Response, error)
}

// ConfigRepository declares the opeations that a Repository needs to support
type ConfigRepository interface {
	AddNewIP(ctx context.Context, label, ip string) error
	RemoveIPRange(ctx context.Context, label string) error
}
