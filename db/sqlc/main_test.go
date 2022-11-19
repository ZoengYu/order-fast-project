package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil{
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connec to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
