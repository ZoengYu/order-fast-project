package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func TestAddMenuItemWithTag(t *testing.T) {
	_, item := CreateRandomMenuItem(t)

	add_ItemTag_arg := CreateMenuItemTagParams{
		ItemID: item.ID,
		ItemTag: util.RandomItemTag(),
	}
	add_ItemTag, err := testQueries.CreateMenuItemTag(context.Background(), add_ItemTag_arg)
	require.NoError(t, err)
	require.Equal(t, add_ItemTag_arg.ItemTag, add_ItemTag.ItemTag)
	get_ItemTag_arg := GetMenuItemTagParams{
		ItemID: item.ID,
		ItemTag: add_ItemTag.ItemTag,
	}
	get_ItemTag, err := testQueries.GetMenuItemTag(context.Background(), get_ItemTag_arg)
	require.NoError(t, err)
	require.NotEmpty(t, get_ItemTag)
	require.Equal(t, get_ItemTag.ID, add_ItemTag.ID)
}

func TestListMenuItemeTag(t *testing.T) {
	_, item := CreateRandomMenuItem(t)
	for i := 0; i < 3; i++ {
		arg := CreateMenuItemTagParams{
			ItemID: item.ID,
			ItemTag: util.RandomItemTag(),
		}
		_, err := testQueries.CreateMenuItemTag(context.Background(), arg)
		require.NoError(t, err)
	}
	ItemTag_list, err := testQueries.ListMenuItemTag(context.Background(), item.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ItemTag_list)
	require.Len(t, ItemTag_list, 3)
}

func TestRemoveMenuItemTag(t *testing.T){
	_, item := CreateRandomMenuItem(t)
	add_ItemTag_arg := CreateMenuItemTagParams{
		ItemID: item.ID,
		ItemTag: util.RandomItemTag(),
	}
	add_ItemTag, err := testQueries.CreateMenuItemTag(context.Background(), add_ItemTag_arg)
	require.NoError(t, err)
	require.NotEmpty(t, add_ItemTag)

	remove_arg := RemoveMenuItemTagParams{
		ItemID: item.ID,
		ItemTag: add_ItemTag.ItemTag,
	}
	err = testQueries.RemoveMenuItemTag(context.Background(), remove_arg)
	require.NoError(t, err)
	get_arg := GetMenuItemTagParams{
		ItemID: item.ID,
		ItemTag: remove_arg.ItemTag,
	}
	get_ItemTag, err := testQueries.GetMenuItemTag(context.Background(), get_arg)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, get_ItemTag)
}
