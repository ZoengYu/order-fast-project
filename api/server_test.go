package api

import (
	"database/sql"
	"net"
	"testing"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func TestServerStartBadAddress(t *testing.T) {
	config, _ := util.LoadConfig(".")

	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	db_service := db.NewDBService(conn)
	config.ServerAddress = "bad_port"

	server, _ := NewServer(config, db_service)
	err := server.Start(config.ServerAddress)
	require.Equal(t, err, err.(*net.OpError))
}
