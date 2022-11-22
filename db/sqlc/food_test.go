package db

import (
	"context"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomMenuFood(t *testing.T) (Menu, MenuFood){
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	arg := AddMenuFoodParams{
		MenuID: menu.ID,
		FoodName: util.RandomFoodName(),
		CustomOption: util.RandomFoodCustom(),
	}
	food, err := testQueries.AddMenuFood(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, food.FoodName, arg.FoodName)
	require.Equal(t, food.CustomOption, arg.CustomOption)
	return menu, food
}

func TestCreateMenuFood(t *testing.T) {
	CreateRandomMenuFood(t)
}

func TestGetMenuFood(t *testing.T) {
	menu, food := CreateRandomMenuFood(t)
	arg := GetMenuFoodParams{
		MenuID: menu.ID,
		FoodName: food.FoodName,
	}
	get_food, err := testQueries.GetMenuFood(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.FoodName, get_food.FoodName)
}
