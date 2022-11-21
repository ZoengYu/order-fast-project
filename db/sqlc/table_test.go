package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTable(t *testing.T) Table{
	store := createRandomStore(t)

	table1_arg := CreateTableParams{
		StoreID: store.ID,
		TableID: 1,
	}
	table, err := testQueries.CreateTable(context.Background(), table1_arg)
	require.NoError(t, err)
	require.Equal(t, table.TableID, table1_arg.TableID)
	require.Equal(t, table.StoreID, store.ID)
	require.Equal(t, table.TableName, table1_arg.TableName)
	return table
}

func getRandomTable(t *testing.T) Table{
	store := getRandomStore(t)
	get_table_arg := GetStoreTableParams{
		StoreID: store.ID,
		TableID: 1,
	}
	table, err := testQueries.GetStoreTable(context.Background(), get_table_arg)
	require.NoError(t, err)
	require.NotEmpty(t, table)

	require.Equal(t, store.ID, table.StoreID)
	return table
}

func TestCreateTable(t *testing.T) {
	createRandomTable(t)
}

func TestGetTable(t *testing.T) {
	getRandomTable(t)
}

func TestUpdateStoreTable(t *testing.T) {
	table := getRandomTable(t)

	arg := UpdateStoreTableParams{
		StoreID: table.StoreID,
		TableID: table.TableID,
		TableName: "earth",
	}
	err := testQueries.UpdateStoreTable(context.Background(), arg)
	require.NoError(t, err)
	updated_table := getRandomTable(t)
	require.NotEmpty(t, updated_table)

	require.Equal(t, "earth", updated_table.TableName)
}

func TestDeleteStoreTable(t *testing.T) {
	table := getRandomTable(t)
	err := testQueries.DeleteStoreTable(context.Background(), table.ID)
	require.NoError(t, err)

	get_table_arg := GetStoreTableParams{
		StoreID: table.StoreID,
		TableID: table.TableID,
	}
	empty_table, err := testQueries.GetStoreTable(context.Background(), get_table_arg)
	require.Empty(t, empty_table)
	require.Equal(t, err, sql.ErrNoRows)

	delRandomStore(t)
}
