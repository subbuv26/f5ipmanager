package http

import (
	"context"
	"encoding/json"
	cfgsvc "f5ipmanager/internal/core/services/configure"
	spec "f5ipmanager/internal/core/spec/configure"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
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
	return kithttp.NewServer(endpoint, decodeAddRequest, encodeResponse)
}

func makeRemoveConfigHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(endpoint, decodeRemoveRequest, encodeResponse)
}

func encodeResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}

func decodeRequest(_ context.Context, request *http.Request, req interface{}) error {

	if request.Body == nil {
		return spec.ErrorInvalidInput
	}

	data, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}

	defer request.Body.Close()

	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	return nil
}

func decodeAddRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.AddIPRangeRequest{}
	err := decodeRequest(ctx, request, req)
	return req, err
}

func decodeRemoveRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := spec.RemoveIPRangeRequest{}
	err := decodeRequest(ctx, request, req)
	return req, err
}
