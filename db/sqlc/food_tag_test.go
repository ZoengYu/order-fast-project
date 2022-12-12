package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func TestAddMenuFoodWithTag(t *testing.T) {
	_, food := CreateRandomMenuFood(t)

	add_foodtag_arg := AddMenuFoodTagParams{
		FoodID: food.ID,
		FoodTag: util.RandomFoodTag(),
	}
	add_foodtag, err := testQueries.AddMenuFoodTag(context.Background(), add_foodtag_arg)
	require.NoError(t, err)
	require.Equal(t, add_foodtag_arg.FoodTag, add_foodtag.FoodTag)
	get_foodtag_arg := GetMenuFoodTagParams{
		FoodID: food.ID,
		FoodTag: add_foodtag.FoodTag,
	}
	get_foodtag, err := testQueries.GetMenuFoodTag(context.Background(), get_foodtag_arg)
	require.NoError(t, err)
	require.NotEmpty(t, get_foodtag)
	require.Equal(t, get_foodtag.ID, add_foodtag.ID)
}

func TestListMenuFoodeTag(t *testing.T) {
	_, food := CreateRandomMenuFood(t)
	for i := 0; i < 3; i++ {
		arg := AddMenuFoodTagParams{
			FoodID: food.ID,
			FoodTag: util.RandomFoodTag(),
		}
		_, err := testQueries.AddMenuFoodTag(context.Background(), arg)
		require.NoError(t, err)
	}
	foodtag_list, err := testQueries.ListMenuFoodTag(context.Background(), food.ID)
	require.NoError(t, err)
	require.NotEmpty(t, foodtag_list)
	require.Len(t, foodtag_list, 3)
}

func TestRemoveMenuFoodTag(t *testing.T){
	menu, food := CreateRandomMenuFood(t)
	add_foodtag_arg := AddMenuFoodTagParams{
		FoodID: food.ID,
		FoodTag: util.RandomFoodTag(),
	}
	add_foodtag, err := testQueries.AddMenuFoodTag(context.Background(), add_foodtag_arg)
	require.NoError(t, err)
	require.NotEmpty(t, add_foodtag)

	remove_arg := RemoveMenuFoodTagParams{
		FoodID: menu.ID,
		FoodTag: add_foodtag.FoodTag,
	}
	err = testQueries.RemoveMenuFoodTag(context.Background(), remove_arg)
	require.NoError(t, err)
	get_arg := GetMenuFoodTagParams{
		FoodID: menu.ID,
		FoodTag: remove_arg.FoodTag,
	}
	get_foodtag, err := testQueries.GetMenuFoodTag(context.Background(), get_arg)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, get_foodtag)
}
