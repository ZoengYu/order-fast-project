package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStore(t *testing.T) Store{
	arg := CreateStoreParams{
		StoreName: util.RandomStoreName(),
		StoreAddress: "address",
		StorePhone: util.RandomPhone(),
		StoreOwner: util.RandomOwner(),
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

func TestCreateStore(t *testing.T) {
	createRandomStore(t)
}

func TestGetStore(t *testing.T) {
	store := createRandomStore(t)
	get_store, err := testQueries.GetStore(context.Background(), store.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get_store)

	require.Equal(t, get_store.StoreName, store.StoreName)
}

func TestGetStoreByName(t *testing.T) {
	store := createRandomStore(t)
	get_store, err := testQueries.GetStoreByName(context.Background(), store.StoreName)
	require.NoError(t, err)
	require.NotEmpty(t, get_store)

	require.Equal(t, get_store.StoreName, store.StoreName)
}

func TestUpdateStore(t *testing.T) {
	store := createRandomStore(t)

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
	store := createRandomStore(t)
	require.NotEmpty(t, store)

	err := testQueries.DeleteStore(context.Background(), store.ID)
	require.NoError(t, err)
	get_store, err := testQueries.GetStoreByName(context.Background(), store.StoreName)
	require.Empty(t, get_store)
	require.Equal(t, err, sql.ErrNoRows)
}
