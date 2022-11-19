package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func getRandomTable(t *testing.T) Table{
	store := getRandomStore(t)
	get_table_arg := GetTableParams{
		StoreID: store.ID,
		TableID: 1,
	}
	table, err := testQueries.GetTable(context.Background(), get_table_arg)
	require.NoError(t, err)
	require.NotEmpty(t, table)

	require.Equal(t, store.ID, table.StoreID)
	return table
}

func TestCreateTable(t *testing.T) {
	store := createRandomStore(t)

	table1_arg := CreateTableParams{
		StoreID: store.ID,
		TableID: 1,
		TableName: "none",
	}
	table, err := testQueries.CreateTable(context.Background(), table1_arg)
	require.NoError(t, err)
	require.Equal(t, table.TableID, table1_arg.TableID)
	require.Equal(t, table.StoreID, store.ID)
}

func TestGetTable(t *testing.T) {
	getRandomTable(t)
}

func TestUpdateTable(t *testing.T) {
	table := getRandomTable(t)

	arg := UpdateTableParams{
		StoreID: table.StoreID,
		TableID: table.TableID,
		TableName: "earth",
	}
	err := testQueries.UpdateTable(context.Background(), arg)
	require.NoError(t, err)
	updated_table := getRandomTable(t)
	require.NotEmpty(t, updated_table)

	require.Equal(t, "earth", updated_table.TableName)
}

func TestDeleteTable(t *testing.T) {
	table := getRandomTable(t)
	err := testQueries.DeleteTable(context.Background(), table.ID)
	require.NoError(t, err)

	get_table_arg := GetTableParams{
		StoreID: table.StoreID,
		TableID: table.TableID,
	}
	empty_table, err := testQueries.GetTable(context.Background(), get_table_arg)
	require.Empty(t, empty_table)
	require.Equal(t, err, sql.ErrNoRows)
	store, _ := testQueries.GetStore(context.Background(), "Harry")
	testQueries.DeleteStore(context.Background(), store.ID)
}
