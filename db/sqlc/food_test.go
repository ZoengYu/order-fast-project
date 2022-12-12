package db

import (
	"context"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomMenuFood(t *testing.T) (Menu, Food){
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	arg := CreateMenuFoodParams{
		MenuID: menu.ID,
		Name: 	util.RandomFoodName(),
	}
	food, err := testQueries.CreateMenuFood(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, food.Name, arg.Name)
	return menu, food
}

func TestCreateMenuFood(t *testing.T) {
	CreateRandomMenuFood(t)
}

func TestGetMenuFood(t *testing.T) {
	menu, food := CreateRandomMenuFood(t)
	arg := GetMenuFoodParams{
		MenuID: menu.ID,
		Name: food.Name,
	}
	get_food, err := testQueries.GetMenuFood(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Name, get_food.Name)
}
