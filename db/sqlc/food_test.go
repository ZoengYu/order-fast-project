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
		Price:  int32(util.RandomInt(20, 100)),
	}
	food, err := testQueries.CreateMenuFood(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, food.Name, arg.Name)
	require.Equal(t, food.Price, arg.Price)
	return menu, food
}

func TestCreateMenuFood(t *testing.T) {
	CreateRandomMenuFood(t)
}

func TestGetMenuFood(t *testing.T) {
	_, food := CreateRandomMenuFood(t)
	get_food, err := testQueries.GetMenuFood(context.Background(), food.ID)
	require.NoError(t, err)
	require.Equal(t, get_food, food)
}
