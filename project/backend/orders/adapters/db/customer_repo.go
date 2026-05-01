package db

import (
	"context"
	"eats/backend/common"
	"eats/backend/common/shared"
	"eats/backend/orders/adapters/db/dbmodels"
	"eats/backend/orders/api/http"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	if db == nil {
		panic("db connection pool cannot be nil")
	}

	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) RegisterCustomer(ctx context.Context, customerUUID common.UUID, customer http.RegisterCustomer) error {
	queries := dbmodels.New(r.db)

	commonAddress, err := openapiAddressToSharedAddress(customer)
	if err != nil {
		return fmt.Errorf("convert address failed: %w", err)
	}

	err = queries.InsertCustomer(ctx, dbmodels.InsertCustomerParams{
		CustomerUuid: customerUUID,
		Name:         customer.Name,
		Email:        string(customer.Email),
		Address:      commonAddress,
		PhoneNumber:  customer.PhoneNumber,
	})
	if err != nil {
		return fmt.Errorf("insert customer failed: %w", err)
	}
	return nil
}

func openapiAddressToSharedAddress(customer http.RegisterCustomer) (shared.Address, error) {
	addr := customer.Address
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
