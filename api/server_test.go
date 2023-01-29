package api

import (
	"database/sql"
	"net"
	"testing"
	"time"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func TestServerStartBadAddress(t *testing.T) {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	db_service := db.NewDBService(conn)
	config.HTTPServerAddress = "bad_port"

	server, _ := NewServer(config, db_service)
	err := server.Start(config.HTTPServerAddress)
	require.Equal(t, err, err.(*net.OpError))
}
