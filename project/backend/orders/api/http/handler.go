package http

import (
	"context"
	"eats/backend/common"
	"eats/backend/common/shared"
	"eats/backend/orders/app"
	"fmt"
)

type Handler struct {
	appService *app.Service
}

func NewHandler(
	appService *app.Service,
) Handler {
	if appService == nil {
		panic("appService cannot be nil")
	}

	return Handler{
		appService: appService,
	}
}

func (h Handler) RegisterCustomer(ctx context.Context, request RegisterCustomerRequestObject) (RegisterCustomerResponseObject, error) {
	commonAddress, err := openapiAddressToSharedAddress(request.Body.Address)
	if err != nil {
		return nil, fmt.Errorf("convert address failed: %w", err)
	}

	customerUUID := common.NewUUIDv7()

	customer := &app.Customer{
		CustomerUUID: customerUUID,
		Name:         request.Body.Name,
		Email:        string(request.Body.Email),
		Address:      commonAddress,
		PhoneNumber:  request.Body.PhoneNumber,
	}

	err = h.appService.RegisterCustomer(ctx, *customer)
	if err != nil {
		return nil, err
	}

	return RegisterCustomer201JSONResponse{
		CustomerUuid: customerUUID,
	}, nil
}

func Register(ctx context.Context, e EchoRouter, handler Handler) error {
	RegisterHandlers(e, NewStrictHandler(handler, nil))

	return nil
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
