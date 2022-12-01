package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStoreMenu(t *testing.T, store Store) Menu{
	arg := CreateStoreMenuParams{
		StoreID: store.ID,
		MenuName: util.RandomMenuName(),
	}
	menu, err := testQueries.CreateStoreMenu(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, arg.MenuName, menu.MenuName)

	require.NotZero(t, menu.ID)
	require.NotZero(t, store.CreatedAt)

	return menu
}

func TestGetRandomMenu(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	get_menu_arg := GetStoreMenuParams{
		StoreID: store.ID,
		ID: menu.ID,
	}
	menu, err := testQueries.GetStoreMenu(context.Background(), get_menu_arg)
	require.NoError(t, err)
	require.NotEmpty(t, menu)
}

func TestCreateStoreMenu(t *testing.T) {
	store := createRandomStore(t)
	createRandomStoreMenu(t, store)
}

func TestUpdateStoreMenu(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	update_menu_arg := UpdateStoreMenuParams{
		StoreID: 	store.ID,
		ID: 		menu.ID,
		MenuName: 	"My Menu2",
	}
	updated_menu, err := testQueries.UpdateStoreMenu(context.Background(), update_menu_arg)
	require.NoError(t, err)

	require.Equal(t, updated_menu.ID, menu.ID)
	require.Equal(t, updated_menu.MenuName, update_menu_arg.MenuName)
}

func TestDeleteMenu(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	err := testQueries.DeleteMenu(context.Background(), menu.ID)
	require.NoError(t, err)
	arg := GetStoreMenuParams{
		StoreID: 	store.ID,
		ID:			menu.ID,
	}
	get_menu, err := testQueries.GetStoreMenu(context.Background(), arg)
	require.Empty(t, get_menu)
	require.Equal(t, err, sql.ErrNoRows)

}
