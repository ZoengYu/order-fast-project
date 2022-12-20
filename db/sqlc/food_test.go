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

func TestGetFood(t *testing.T) {
	_, food := CreateRandomMenuFood(t)
	get_food, err := testQueries.GetFood(context.Background(), food.ID)
	require.NoError(t, err)
	require.Equal(t, get_food, food)
}

func TestDeleteFood(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	// create two food in the menu
	arg1 := CreateMenuFoodParams{
		MenuID: menu.ID,
		Name: 	util.RandomFoodName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	food, err := testQueries.CreateMenuFood(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, food)
	arg2 := CreateMenuFoodParams{
		MenuID: menu.ID,
		Name: 	util.RandomFoodName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	food2, err := testQueries.CreateMenuFood(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, food2)
	// delete one food
	del_arg := DeleteMenuFoodParams{
		ID:		food.ID,
		MenuID: menu.ID,
	}
	err = testQueries.DeleteMenuFood(context.Background(), del_arg)
	require.NoError(t, err)
	food_list, err := testQueries.ListMenuFood(context.Background(), menu.ID)
	require.NoError(t, err)
	require.NotContains(t, food_list, food)
	require.Contains(t, food_list, food2)
}

func TestDeleteFoodAll(t *testing.T) {
	store := createRandomStore(t)
	menu := createRandomStoreMenu(t, store)
	arg := CreateMenuFoodParams{
		MenuID: menu.ID,
		Name: 	util.RandomFoodName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateMenuFood(context.Background(), arg)
		require.NoError(t, err)
	}
	food_list, err := testQueries.ListMenuFood(context.Background(), menu.ID)
	require.NoError(t, err)
	require.Len(t, food_list, 3)
	err = testQueries.DeleteMenuFoodAll(context.Background(), menu.ID)
	require.NoError(t, err)
	food_list, err = testQueries.ListMenuFood(context.Background(), menu.ID)
	require.NoError(t, err)
	require.Len(t, food_list, 0)
}
