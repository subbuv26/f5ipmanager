package configure

import (
	"context"
	ports "f5ipmanager/internal/core/ports/configure"
	spec "f5ipmanager/internal/core/spec/configure"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddIPRange    endpoint.Endpoint
	RemoveIPRange endpoint.Endpoint
}

func MakeEndpoints(svc ports.ConfigService) Endpoints {
	return Endpoints{
		AddIPRange:    makeAddConfigEndpoint(svc),
		RemoveIPRange: makeRemoveConfigEndpoint(svc),
	}
}

func makeAddConfigEndpoint(svc ports.ConfigService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(spec.AddIPRangeRequest)
		return svc.AddIPRange(ctx, req)
	}
}

func makeRemoveConfigEndpoint(svc ports.ConfigService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(spec.RemoveIPRangeRequest)
		return svc.RemoveIPRange(ctx, req)
	}
}
