package http

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	cfgsvc "f5ipmanager/internal/core/services/manager"
	spec "f5ipmanager/internal/core/spec/manager"
	httputils "f5ipmanager/pkg/utils/http"
)

type HTTPHandlers struct {
	AllocateIPAddress   http.Handler
	DeallocateIPAddress http.Handler
}

func GetHTTPHandlers(endpoints cfgsvc.Endpoints) HTTPHandlers {
	return HTTPHandlers{
		AllocateIPAddress:   makeAllocateIPAddressHTTPHandler(endpoints.AllocateIPAddress),
		DeallocateIPAddress: makeDeallocateIPAddressHTTPHandler(endpoints.DeallocateIPAddress),
	}
}

func makeAllocateIPAddressHTTPHandler(ep endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(ep, decodeAllocateRequest, httputils.EncodeResponse)
}

func makeDeallocateIPAddressHTTPHandler(ep endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(ep, decodeDeallocateRequest, httputils.EncodeResponse)
}

func decodeAllocateRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.AllocateRequest{}
	err := httputils.DecodeRequest(ctx, request, req)
	return req, err
}

func decodeDeallocateRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.DeallocateRequest{}
	err := httputils.DecodeRequest(ctx, request, req)
	return req, err
}
