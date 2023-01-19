package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStore(t *testing.T) Store {
	user := createRandomUser(t)
	arg := CreateStoreParams{
		Owner:   user.Username,
		Name:    util.RandomName(),
		Address: "address",
		Phone:   util.RandomPhone(),
		Manager: "王小棣s",
	}
	store, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, user.Username, store.Owner)
	require.Equal(t, arg.Name, store.Name)
	require.Equal(t, arg.Address, store.Address)
	require.Equal(t, arg.Phone, store.Phone)
	require.Equal(t, arg.Manager, store.Manager)

	require.NotZero(t, store.ID)
	require.NotZero(t, store.CreatedAt)

	return store
}

func createMultipleStore(t *testing.T, num int) []Store {
	var storeList []Store
	for i := 0; i < num; i++ {
		store := createRandomStore(t)
		storeList = append(storeList, store)
	}
	return storeList
}

func TestCreateStore(t *testing.T) {
	createRandomStore(t)
}

func TestGetStore(t *testing.T) {
	store := createRandomStore(t)
	get_store, err := testQueries.GetStore(context.Background(), store.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get_store)

	require.Equal(t, get_store.Name, store.Name)
}

func TestGetStoreByName(t *testing.T) {
	stores := createMultipleStore(t, 3)
	arg := ListStoresByNameParams{
		Name:   stores[0].Name,
		Limit:  3,
		Offset: 0,
	}
	get_stores, err := testQueries.ListStoresByName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, get_stores)

	require.Equal(t, get_stores[0], stores[0])
}

func TestUpdateStore(t *testing.T) {
	store := createRandomStore(t)

	arg := UpdateStoreParams{
		ID:      store.ID,
		Owner:   store.Owner,
		Name:    store.Name,
		Address: store.Address,
		Phone:   store.Phone,
		Manager: "Alex",
	}
	new_store, err := testQueries.UpdateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, new_store)

	require.Equal(t, "Alex", new_store.Manager)
}

func TestDeleteStore(t *testing.T) {
	store := createRandomStore(t)
	require.NotEmpty(t, store)

	err := testQueries.DeleteStore(context.Background(), store.ID)
	require.NoError(t, err)
	get_store, err := testQueries.GetStore(context.Background(), store.ID)
	require.Empty(t, get_store)
	require.Equal(t, err, sql.ErrNoRows)
}
