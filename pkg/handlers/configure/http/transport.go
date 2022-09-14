package http

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	cfgsvc "f5ipmanager/internal/core/services/configure"
	spec "f5ipmanager/internal/core/spec/configure"
	httputils "f5ipmanager/pkg/utils/http"
)

type HTTPHandlers struct {
	AddIPRange    http.Handler
	RemoveIPRange http.Handler
}

func GetHTTPHandlers(endpoints cfgsvc.Endpoints) HTTPHandlers {
	return HTTPHandlers{
		AddIPRange:    makeAddConfigHTTPHandler(endpoints.AddIPRange),
		RemoveIPRange: makeRemoveConfigHTTPHandler(endpoints.RemoveIPRange),
	}
}

func makeAddConfigHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(endpoint, decodeAddRequest, httputils.EncodeResponse)
}

func makeRemoveConfigHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(endpoint, decodeRemoveRequest, httputils.EncodeResponse)
}

func decodeAddRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.AddIPRangeRequest{}
	err := httputils.DecodeRequest(ctx, request, req)
	return req, err
}

func decodeRemoveRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.RemoveIPRangeRequest{}
	err := httputils.DecodeRequest(ctx, request, req)
	return req, err
}
