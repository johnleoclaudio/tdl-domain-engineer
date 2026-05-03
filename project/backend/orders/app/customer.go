package app

import (
	"context"
	"strings"

	"eats/backend/common"
	"eats/backend/common/shared"
)

type CustomerUUID struct {
	common.UUID
}

type Customer struct {
	CustomerUUID CustomerUUID
	Name         string
	Email        string
	Address      shared.Address
	PhoneNumber  string
}

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, customer Customer) error
}

func (s *Service) RegisterCustomer(ctx context.Context, customer Customer) error {
	validationErrors := []common.ErrorDetails{}

	if customer.CustomerUUID.IsZero() {
		errDetails := common.ErrorDetails{
			EntityType: "customer",
			ErrorSlug:  "empty-uuid",
		}

		validationErrors = append(validationErrors, errDetails)
	}

	name := strings.TrimSpace(customer.Name)
	if name == "" {
		errDetails := common.ErrorDetails{
			EntityType: "customer",
			ErrorSlug:  "empty-name",
		}

		validationErrors = append(validationErrors, errDetails)
	}

	email := strings.TrimSpace(customer.Email)
	if email == "" {
		errDetails := common.ErrorDetails{
			EntityType: "customer",
			ErrorSlug:  "empty-email",
		}

		validationErrors = append(validationErrors, errDetails)
	}

	if customer.Address.IsZero() {
		errDetails := common.ErrorDetails{
			EntityType: "customer",
			ErrorSlug:  "empty-address",
		}

		validationErrors = append(validationErrors, errDetails)
	}

	phoneNumber := strings.TrimSpace(customer.PhoneNumber)
	if phoneNumber == "" {
		errDetails := common.ErrorDetails{
			EntityType: "customer",
			ErrorSlug:  "empty-phone-number",
		}

		validationErrors = append(validationErrors, errDetails)
	}

	if len(validationErrors) != 0 {
		return common.NewInvalidInputError("invalid_customer_data", "Invalid customer data").WithDetails(validationErrors)
	}

	return s.customerRepository.RegisterCustomer(ctx, customer)
}
