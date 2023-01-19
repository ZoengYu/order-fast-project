package db

import (
	"context"
	"database/sql"
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

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	get_account, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, account, get_account)
}

func TestDelAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	_, err = testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	updated_arg := UpdateAccountParams{
		ID:    account.ID,
		Owner: "Updated_Owner",
	}
	updated_account, err := testQueries.UpdateAccount(context.Background(), updated_arg)
	require.NoError(t, err)
	require.Equal(t, updated_account.Owner, updated_arg.Owner)
}
