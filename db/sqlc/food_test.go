package db

import (
	"context"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomMenuItem(t *testing.T) (Menu, Item){
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	arg := CreateMenuItemParams{
		MenuID: menu.ID,
		Name: 	util.RandomItemName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	item, err := testQueries.CreateMenuItem(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, item.Name, arg.Name)
	require.Equal(t, item.Price, arg.Price)
	return menu, item
}

func TestCreateMenuItem(t *testing.T) {
	CreateRandomMenuItem(t)
}

func TestGetItem(t *testing.T) {
	_, item := CreateRandomMenuItem(t)
	get_item, err := testQueries.GetItem(context.Background(), item.ID)
	require.NoError(t, err)
	require.Equal(t, get_item, item)
}

func TestDeleteItem(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	// create two item in the menu
	arg1 := CreateMenuItemParams{
		MenuID: menu.ID,
		Name: 	util.RandomItemName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	item, err := testQueries.CreateMenuItem(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, item)
	arg2 := CreateMenuItemParams{
		MenuID: menu.ID,
		Name: 	util.RandomItemName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	item2, err := testQueries.CreateMenuItem(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, item2)
	// delete one item
	del_arg := DeleteMenuItemParams{
		ID:		item.ID,
		MenuID: menu.ID,
	}
	err = testQueries.DeleteMenuItem(context.Background(), del_arg)
	require.NoError(t, err)
	item_list, err := testQueries.ListAllMenuItem(context.Background(), menu.ID)
	require.NoError(t, err)
	require.NotContains(t, item_list, item)
	require.Contains(t, item_list, item2)
}

func TestDeleteItemAll(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	arg := CreateMenuItemParams{
		MenuID: menu.ID,
		Name: 	util.RandomItemName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateMenuItem(context.Background(), arg)
		require.NoError(t, err)
	}
	item_list, err := testQueries.ListAllMenuItem(context.Background(), menu.ID)
	require.NoError(t, err)
	require.Len(t, item_list, 3)
	err = testQueries.DeleteMenuItemAll(context.Background(), menu.ID)
	require.NoError(t, err)
	item_list, err = testQueries.ListAllMenuItem(context.Background(), menu.ID)
	require.NoError(t, err)
	require.Len(t, item_list, 0)
}
