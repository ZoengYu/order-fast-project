package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/gapi"
	"github.com/ZoengYu/order-fast-project/pb"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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
	go runGatewayServer(config, db_service)
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
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGatewayServer(config util.Config, db_service db.DBService) {
	gapi_server, err := gapi.NewServer(config, db_service)
	if err != nil {
		log.Fatal("cannot create the gapi server:", err)
	}

	// convert from camel case to snake case to align with the format specified in proto
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterOrderFastHandlerServer(ctx, grpcMux, gapi_server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// server the swagger UI for gRPC server
	fileServer := http.FileServer(http.Dir("./docs/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger", fileServer))

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP server:", err)
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
