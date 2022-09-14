package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func EncodeResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}

func DecodeRequest(_ context.Context, request *http.Request, req interface{}) error {

	if request.Body == nil {
		return errors.New("empty body")
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
