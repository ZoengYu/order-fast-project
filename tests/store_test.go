package db

import (
	"context"
	"testing"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestCreateStore(t *testing.T) {
	arg := db.CreateStoreParams{
		StoreName: "Harry",
		StoreAddress: "address",
		StorePhone: "0900000000",
		StoreOwner: "王震",
		StoreManager: "王小棣s",
	}

	store, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, arg.StoreName, store.StoreName)
	require.Equal(t, arg.StoreAddress, store.StoreAddress)
	require.Equal(t, arg.StorePhone, store.StorePhone)
	require.Equal(t, arg.StoreOwner, store.StoreOwner)
	require.Equal(t, arg.StoreManager, store.StoreManager)

	require.NotZero(t, store.ID)
	require.NotZero(t, store.CreatedAt)
}
