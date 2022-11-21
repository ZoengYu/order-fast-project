package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomStore(t *testing.T) Store{
	arg := CreateStoreParams{
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

	return store
}

func getRandomStore(t *testing.T) Store{
	store, err := testQueries.GetStore(context.Background(), "Harry")
	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, "Harry", store.StoreName)
	return store
}

func delRandomStore(t *testing.T) {
	store, _ := testQueries.GetStore(context.Background(), "Harry")
	require.NotEmpty(t, store)

	err := testQueries.DeleteStore(context.Background(), store.ID)
	require.NoError(t, err)
	get_store, err := testQueries.GetStore(context.Background(), "Harry")
	require.Empty(t, get_store)
	require.Equal(t, err, sql.ErrNoRows)
}
func TestCreateStore(t *testing.T) {
	createRandomStore(t)
}

func TestGetStore(t *testing.T) {
	getRandomStore(t)
}

func TestUpdateStore(t *testing.T) {
	store, _ := testQueries.GetStore(context.Background(), "Harry")
	require.NotEmpty(t, store)

	arg := UpdateStoreParams{
		ID: store.ID,
		StoreName: store.StoreName,
		StoreAddress: store.StoreAddress,
		StorePhone: store.StorePhone,
		StoreOwner: store.StoreOwner,
		StoreManager: "Alex",
	}
	new_store, err := testQueries.UpdateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, new_store)

	require.Equal(t, "Alex", new_store.StoreManager)
}

func TestDeleteStore(t *testing.T) {
	delRandomStore(t)
}
