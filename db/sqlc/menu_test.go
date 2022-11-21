package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMenu(t *testing.T) Menu{
	store := createRandomStore(t)
	arg := CreateStoreMenuParams{
		StoreID: store.ID,
		MenuName: "My Menu",
	}
	menu, err := testQueries.CreateStoreMenu(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, arg.MenuName, menu.MenuName)

	require.NotZero(t, menu.ID)
	require.NotZero(t, store.CreatedAt)

	return menu
}

func getRandomMenu(t *testing.T) Menu{
	store := getRandomStore(t)
	get_menu_arg := GetStoreMenuParams{
		StoreID: store.ID,
		MenuName: "My Menu",
	}
	menu, err := testQueries.GetStoreMenu(context.Background(), get_menu_arg)
	require.NoError(t, err)
	require.NotEmpty(t, menu)
	return menu
}
func TestCreateStoreMenu(t *testing.T) {
	createRandomMenu(t)
}

func TestUpdateStoreMenu(t *testing.T) {
	menu := getRandomMenu(t)
	update_menu_arg := UpdateStoreMenuParams{
		ID: menu.ID,
		MenuName: "My Menu2",
	}
	updated_menu, err := testQueries.UpdateStoreMenu(context.Background(), update_menu_arg)
	require.NoError(t, err)

	require.Equal(t, updated_menu.ID, menu.ID)
	require.Equal(t, updated_menu.MenuName, update_menu_arg.MenuName)
}

func TestAddMenuFoodWithTag(t *testing.T) {
	store := getRandomStore(t)
	get_menu_arg := GetStoreMenuParams{
		StoreID: store.ID,
		MenuName: "My Menu2",
	}
	menu, _ := testQueries.GetStoreMenu(context.Background(), get_menu_arg)
	add_food_arg := AddMenuFoodParams{
		MenuID: menu.ID,
		FoodName: "food1",
		CustomOption: []string{"小麥麵包 +NT5", "去醬", "雙煎蛋 +NT10"},
	}
	menu_food, err := testQueries.AddMenuFood(context.Background(), add_food_arg)
	require.NoError(t, err)
	require.NotEmpty(t, menu_food)
	require.Equal(t, menu_food.FoodName, add_food_arg.FoodName)

	add_foodtag_arg := AddMenuFoodTagParams{
		MenuFoodID: menu_food.ID,
		FoodTag: "三明治",
	}
	add_foodtag, err := testQueries.AddMenuFoodTag(context.Background(), add_foodtag_arg)
	require.NoError(t, err)
	require.Equal(t, add_foodtag_arg.FoodTag, add_foodtag.FoodTag)
	add_foodtag_arg2 := AddMenuFoodTagParams{
		MenuFoodID: menu_food.ID,
		FoodTag: "top1",
	}
	add_foodtag2, _ := testQueries.AddMenuFoodTag(context.Background(), add_foodtag_arg2)
	foodtag_list, err := testQueries.ListMenuFoodTag(context.Background(), menu_food.ID)
	require.NoError(t, err)
	require.Contains(t, foodtag_list, add_foodtag.FoodTag, add_foodtag2.FoodTag)
}

func TestRemoveMenuFoodTag(t *testing.T){
	store := getRandomStore(t)
	get_menu_arg := GetStoreMenuParams{
		StoreID: store.ID,
		MenuName: "My Menu2",
	}
	menu, _ := testQueries.GetStoreMenu(context.Background(), get_menu_arg)
	remove_foodtag_arg := RemoveMenuFoodTagParams{
		MenuFoodID: menu.ID,
		FoodTag: "三明治",
	}
	err := testQueries.RemoveMenuFoodTag(context.Background(), remove_foodtag_arg)
	require.NoError(t, err)
	get_menu_food_arg := GetMenuFoodParams{
		MenuID: menu.ID,
		FoodName: "food",
	}
	menu_food, _ := testQueries.GetMenuFood(context.Background(), get_menu_food_arg)
	foodtag_list, err := testQueries.ListMenuFoodTag(context.Background(), menu_food.ID)
	require.NoError(t, err)
	require.NotContains(t, foodtag_list, remove_foodtag_arg.FoodTag)

	remove_foodtag_arg2 := RemoveMenuFoodTagParams{
		MenuFoodID: menu.ID,
		FoodTag: "top1",
	}
	err = testQueries.RemoveMenuFoodTag(context.Background(), remove_foodtag_arg2)
	require.NoError(t, err)
	foodtag_list, _ = testQueries.ListMenuFoodTag(context.Background(), menu_food.ID)
	require.Empty(t, foodtag_list)
	delRandomStore(t)
}
