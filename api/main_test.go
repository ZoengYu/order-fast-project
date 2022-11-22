package api

import (
	"os"
	"testing"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, mockdb db.DBService) *Server{
	config := util.Config{}

	server, err := NewServer(config, mockdb)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
