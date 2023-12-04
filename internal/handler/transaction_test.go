package handler

import (
	"context"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

type trServiceMock struct {
	create func(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error)
	get    func(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error)
}

func (h *trServiceMock) Create(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	return h.create(ctx, req)
}

func (h *trServiceMock) Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
	return h.get(ctx, req)
}

func TestTransactionHandler_Create(t *testing.T) {
	resp := api.CreateTransactionResponse{
		ID: mocks.OptString("123-123"),
	}
	count := 0
	tsMock := trServiceMock{
		create: func(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
			count++
			return &resp, nil
		},
	}
	r := api.CreateTransactionRequest{}
	th := transactionHandler{
		&tsMock,
	}

	actualResp, err := th.Create(context.Background(), &r)
	assert.Nil(t, err)
	assert.Equal(t, &resp, actualResp, "Create should pass back the object received from the transaction service unmodified")
}

func TestTransactionHandler_Get(t *testing.T) {
	resp := api.GetTransactionResponse{
		ID: mocks.OptString("123-123"),
	}
	count := 0
	tsMock := trServiceMock{
		get: func(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
			count++
			return &resp, nil
		},
	}
	r := api.GetTransactionRequest{}
	th := transactionHandler{
		&tsMock,
	}

	actualResp, err := th.Get(context.Background(), &r)
	assert.Nil(t, err)
	assert.Equal(t, &resp, actualResp, "Get should pass back the object received from the transaction service unmodified")
}
