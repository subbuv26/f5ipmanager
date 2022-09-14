package manager

import (
	"context"
	ports "f5ipmanager/internal/core/ports/manager"
	spec "f5ipmanager/internal/core/spec/manager"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AllocateIPAddress   endpoint.Endpoint
	DeallocateIPAddress endpoint.Endpoint
}

func MakeEndpoints(svc ports.ManagerService) Endpoints {
	return Endpoints{
		AllocateIPAddress:   makeAllocateIPAddressEndpoint(svc),
		DeallocateIPAddress: makeDeallocateIPAddressEndpoint(svc),
	}
}

func makeAllocateIPAddressEndpoint(svc ports.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(spec.AllocateRequest)

		return svc.AllocateIPAddress(ctx, req)
	}
}

func makeDeallocateIPAddressEndpoint(svc ports.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(spec.DeallocateRequest)
		resp, err := svc.DeallocateIPAddress(ctx, req)
		if err != nil {
			return false, err
		}

		return resp.Success, nil
	}
}
