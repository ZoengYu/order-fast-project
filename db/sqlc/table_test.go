package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStoreTable(t *testing.T, store Store) Table{

	table1_arg := CreateTableParams{
		StoreID: store.ID,
		TableID: util.RandomTableNumber(),
	}
	table, err := testQueries.CreateTable(context.Background(), table1_arg)
	require.NoError(t, err)
	require.Equal(t, table.TableID, table1_arg.TableID)
	require.Equal(t, table.StoreID, store.ID)
	require.Equal(t, table.TableName, table1_arg.TableName)
	return table
}

func TestCreateStoreTable(t *testing.T) {
	store := createRandomStore(t)
	createRandomStoreTable(t, store)
}

func TestGetTable(t *testing.T) {
	store := createRandomStore(t)
	table := createRandomStoreTable(t, store)
	get_table_arg := GetStoreTableParams{
		StoreID: store.ID,
		TableID: table.TableID,
	}
	table, err := testQueries.GetStoreTable(context.Background(), get_table_arg)
	require.NoError(t, err)
	require.NotEmpty(t, table)

	require.Equal(t, store.ID, table.StoreID)
}

func TestUpdateStoreTable(t *testing.T) {
	store := createRandomStore(t)
	table := createRandomStoreTable(t, store)

	arg := UpdateStoreTableParams{
		StoreID: table.StoreID,
		TableID: table.TableID,
		TableName: "earth",
	}
	err := testQueries.UpdateStoreTable(context.Background(), arg)
	require.NoError(t, err)

	get_table_arg := GetStoreTableParams{
		StoreID: store.ID,
		TableID: table.TableID,
	}
	table, err = testQueries.GetStoreTable(context.Background(), get_table_arg)
	require.NoError(t, err)

	require.Equal(t, "earth", table.TableName)
}

func TestDeleteStoreTable(t *testing.T) {
	store := createRandomStore(t)
	table := createRandomStoreTable(t, store)
	err := testQueries.DeleteStoreTable(context.Background(), table.ID)
	require.NoError(t, err)

	get_table_arg := GetStoreTableParams{
		StoreID: store.ID,
		TableID: table.TableID,
	}
	empty_table, err := testQueries.GetStoreTable(context.Background(), get_table_arg)
	require.Empty(t, empty_table)
	require.Equal(t, err, sql.ErrNoRows)
}

func TestListStoreTable(t *testing.T) {
	store := createRandomStore(t)
	for i := 0; i < 5; i++ {
		createRandomStoreTable(t, store)
	}

	arg := ListStoreTablesParams{
		StoreID: store.ID,
		Limit: 5,
		Offset: 1,
	}

	tables, err := testQueries.ListStoreTables(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tables, 4)
}
