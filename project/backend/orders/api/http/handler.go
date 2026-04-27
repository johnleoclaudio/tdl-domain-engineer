package http

import (
	"context"
	"eats/backend/common"
	"eats/backend/common/shared"
	"eats/backend/orders/adapters/db/dbmodels"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	pgxDb *pgxpool.Pool
}

func NewHandler(pgxDb *pgxpool.Pool) Handler {
	if pgxDb == nil {
		panic("db cannot be nil")
	}
	return Handler{
		pgxDb: pgxDb,
	}
}

func openapiAddressToSharedAddress(addr Address) (shared.Address, error) {
	sharedAddr, err := shared.NewAddress(
		addr.Line1,
		addr.Line2,
		addr.PostalCode,
		addr.City,
		addr.CountryCode,
	)
	if err != nil {
		return shared.Address{}, err
	}
	return sharedAddr, nil
}

func (h Handler) RegisterCustomer(ctx context.Context, request RegisterCustomerRequestObject) (RegisterCustomerResponseObject, error) {
	customer := request.Body
	customerUUID := common.NewUUIDv7()

	queries := dbmodels.New(h.pgxDb)

	commonAddress, err := openapiAddressToSharedAddress(customer.Address)
	if err != nil {
		return nil, fmt.Errorf("convert address failed: %w", err)
	}

	err = queries.InsertCustomer(ctx, dbmodels.InsertCustomerParams{
		CustomerUuid: customerUUID,
		Name:         customer.Name,
		Email:        string(customer.Email),
		Address:      commonAddress,
		PhoneNumber:  customer.PhoneNumber,
	})
	if err != nil {
		slog.Error("Failed to Insert", err)
		return RegisterCustomer400JSONResponse{}, nil
	}

	return RegisterCustomer201JSONResponse{
		CustomerUuid: customerUUID,
	}, nil
}

func Register(ctx context.Context, e EchoRouter, handler Handler) error {
	RegisterHandlers(e, NewStrictHandler(handler, nil))

	return nil
}
