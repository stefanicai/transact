package transaction

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/mocks"
	"github.com/stefanicai/transact/internal/model"
	"github.com/stefanicai/transact/internal/persistence/inmem"
	"math/big"
	"time"

	. "github.com/onsi/gomega"
)

type forexServiceMock struct {
	convert func(countryName string, amount big.Rat) (*big.Rat, error)
}

func (f *forexServiceMock) Convert(countryName string, amount big.Rat) (*big.Rat, error) {
	return f.convert(countryName, amount)
}

var _ = Describe("TransactionService", func() {
	expectedDate := time.Date(2021, 04, 01, 0, 0, 0, 0, time.UTC)
	When("Create transaction", func() {
		It("successfully stores it in the in memory store", func() {
			ctx := context.Background()
			dao := inmem.MakeTransactionDao()
			forexMock := forexServiceMock{
				convert: func(countryName string, amount big.Rat) (*big.Rat, error) {
					return &big.Rat{}, nil
				},
			}
			s := NewService(dao, &forexMock)
			resp, err := s.Create(ctx, &api.CreateTransactionRequest{
				Description: api.OptString{Value: "test transaction", Set: true},
				Amount:      api.OptString{Value: "10.32", Set: true},
				Date:        api.OptString{Value: "2021-04-01", Set: true},
			})
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
			tr, err := dao.Get(ctx, resp.ID.Value)
			Expect(err).To(BeNil(), "transaction should be in the store")
			Expect(tr.Date).To(Equal(expectedDate))
		})

		It("Fails when request is invalid - amount not provided", func() {
			ctx := context.Background()
			dao := inmem.MakeTransactionDao()
			forexMock := forexServiceMock{
				convert: func(countryName string, amount big.Rat) (*big.Rat, error) {
					return &big.Rat{}, nil
				},
			}
			s := NewService(dao, &forexMock)
			resp, err := s.Create(ctx, &api.CreateTransactionRequest{
				Description: api.OptString{Value: "test transaction", Set: true},
				Date:        api.OptString{Value: "2021-04-01", Set: true},
			})
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})

	When("Get transaction", func() {
		id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
		amount := big.NewRat(10, 3)
		tr := model.Transaction{
			ID:          id,
			Description: "some description",
			Date:        expectedDate,
			AmountInUSD: *amount,
		}

		It("succeeds when transaction is in the store", func() {
			ctx := context.Background()
			dao := inmem.MakeTransactionDao()
			Expect(dao.Store(ctx, &tr)).To(BeNil())
			forexMock := forexServiceMock{
				convert: func(countryName string, amount big.Rat) (*big.Rat, error) {
					return &big.Rat{}, nil
				},
			}
			s := NewService(dao, &forexMock)
			resp, err := s.Get(ctx, &api.GetTransactionRequest{
				ID:      mocks.OptString(id),
				Country: mocks.OptString("Australia"),
			})
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("fails when transaction is not the store", func() {
			ctx := context.Background()
			dao := inmem.MakeTransactionDao()
			forexMock := forexServiceMock{
				convert: func(countryName string, amount big.Rat) (*big.Rat, error) {
					return &big.Rat{}, nil
				},
			}
			s := NewService(dao, &forexMock)
			resp, err := s.Get(ctx, &api.GetTransactionRequest{
				ID:      mocks.OptString(id),
				Country: mocks.OptString("Australia"),
			})
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})
})
