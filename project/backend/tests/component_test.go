package tests_test

import (
	"context"
	"eats/backend/common/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponent_CriticalFlow(t *testing.T) {
	t.Parallel()

	testClients := newTestClients(t)

	t.Run("TODO", func(t *testing.T) {
		randomCountry := testutils.GenerateRandomCountry()
		customerUUID := registerCustomerInCity(context.TODO(), t, testClients, randomCountry, "Some city")
		assert.NotEmpty(t, customerUUID)
	})
}
