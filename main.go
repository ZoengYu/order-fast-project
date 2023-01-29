package main

import (
	"database/sql"
	"log"
	"net"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/gapi"
	"github.com/ZoengYu/order-fast-project/pb"
	util "github.com/ZoengYu/order-fast-project/utils"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db_service := db.NewDBService(conn)
	runGrpcServer(config, db_service)

}

func runGrpcServer(config util.Config, db_service db.DBService) {
	gapi_server, err := gapi.NewServer(config, db_service)
	if err != nil {
		log.Fatal("cannot create the gapi server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderFastServer(grpcServer, gapi_server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

// func runGinServer(config util.Config, db_service db.DBService) {
// 	api_server, err := api.NewServer(config, db_service)
// 	if err != nil {
// 		log.Fatal("cannot create the api server:", err)
// 	}

// 	err = api_server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot start the api server:", err)
// 	}
// }
