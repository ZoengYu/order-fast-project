package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connec to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}

func CleanUpDBData() error {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connec to db:", err)
	}
	rows, err := db.Query(
		"SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_name != 'schema_migrations'")
	if err != nil {
		log.Fatal("cannot query table from db:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
		if err != nil {
			return err
		}
	}
	return nil
}
