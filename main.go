package main

import (
	"database/sql"
	"log"

	"github.com/ZoengYu/order-fast-project/api"
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	_ "github.com/lib/pq"
)

func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db_service := db.NewDBService(conn)

	api_server, err := api.NewServer(config, db_service)
	if err != nil {
		log.Fatal("cannot create the api server:", err)
	}

	err = api_server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the api server:", err)
	}
}
