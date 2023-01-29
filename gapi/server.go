package gapi

import (
	"fmt"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/pb"
	"github.com/ZoengYu/order-fast-project/token"
	util "github.com/ZoengYu/order-fast-project/utils"
)

// Server serves gRPC requests for our order-fast service.
type Server struct {
	pb.UnimplementedOrderFastServer
	db_service db.DBService
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, db_service db.DBService) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker :%w", err)
	}
	server := &Server{
		db_service: db_service,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
