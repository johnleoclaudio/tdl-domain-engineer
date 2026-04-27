package http

import (
	"context"
	"eats/backend/common"
	"eats/backend/common/shared"
	"eats/backend/orders/adapters/db/dbmodels"
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

func (h Handler) RegisterCustomer(ctx context.Context, request RegisterCustomerRequestObject) (RegisterCustomerResponseObject, error) {
	customerUUID := common.NewUUIDv7()

	arg := dbmodels.InsertCustomerParams{
		CustomerUuid: customerUUID,
		Email:        string(request.Body.Email),
		Address: shared.Address{
			Line1:       request.Body.Address.Line1,
			Line2:       request.Body.Address.Line2,
			PostalCode:  request.Body.Address.PostalCode,
			City:        request.Body.Address.City,
			CountryCode: request.Body.Address.CountryCode,
		},
	}

	slog.Info("TANGINA", arg)

	queries := dbmodels.New(h.pgxDb)
	queries.InsertCustomer(ctx, arg)
	// if err != nil {
	//
	// 	return RegisterCustomer400JSONResponse{}, nil
	// }

	return RegisterCustomer201JSONResponse{
		CustomerUuid: customerUUID,
	}, nil
}

func Register(ctx context.Context, e EchoRouter, handler Handler) error {
	RegisterHandlers(e, NewStrictHandler(handler, nil))

	return nil
}
