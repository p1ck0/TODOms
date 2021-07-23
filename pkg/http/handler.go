package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/p1ck0/TODOms/pkg/endpoints"
	"github.com/p1ck0/TODOms/pkg/service"
)

func MakeHTTPHandler(s service.Serv, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoints.MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/todo").Handler(httptransport.NewServer(
		e.Create,
		decodeCreateTODORequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/todo").Handler(httptransport.NewServer(
		e.Get,
		decodeGetTODORequest,
		encodeResponse,
		options...,
	))

	// r.Methods("PUT").Path("/todo/settimeout").Handler(httptransport.NewServer(
	// 	e.SetTimeOut,
	// 	decodeSetTimeoutRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	return r
}

func decodeCreateTODORequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetTODORequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetRequest
	return req, nil
}

// func decodeSetTimeoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req endpoints.SetTimeOutRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrNotFound:
		return http.StatusNotFound
	case service.ErrAlreadyExists, service.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
