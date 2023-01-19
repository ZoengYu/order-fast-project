package db

import (
	"context"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	account, err := testQueries.CreateAccount(context.Background(), util.RandomOwner())
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
