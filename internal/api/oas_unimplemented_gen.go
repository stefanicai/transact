// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Create implements create operation.
//
// Create a new transaction.
//
// POST /create
func (UnimplementedHandler) Create(ctx context.Context, req *CreateTransactionRequest) (r *CreateTransactionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// Get implements get operation.
//
// Get a new transaction in a specific currency.
//
// GET /get
func (UnimplementedHandler) Get(ctx context.Context, req *GetTransactionRequest) (r *GetTransactionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}