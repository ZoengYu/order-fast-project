package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil{
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connec to db:", err)
	}

	testQueries = db.New(conn)
	os.Exit(m.Run())
}
