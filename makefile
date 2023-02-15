ORDERFAST_CONTAINER := order_fast_db
ORDERFAST_DB := order_fast
ORDERFAST_USER := runner
GOFMT := gofmt
postgres:
	docker run --name ${ORDERFAST_CONTAINER} -p 5432:5432 -e POSTGRES_USER=runner -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it ${ORDERFAST_CONTAINER} createdb --username=runner --owner=runner ${ORDERFAST_DB}

dropdb:
	docker exec -it ${ORDERFAST_CONTAINER} dropdb -U ${ORDERFAST_USER} ${ORDERFAST_DB}

migrateup:
	migrate -path db/migration -database "postgresql://${ORDERFAST_USER}:password@localhost:5432/${ORDERFAST_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${ORDERFAST_USER}:password@localhost:5432/${ORDERFAST_DB}?sslmode=disable" -verbose down -all

migrateforce:
	migrate -path db/migration -database "postgresql://${ORDERFAST_USER}:password@localhost:5432/${ORDERFAST_DB}?sslmode=disable" force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen --package mockdb --destination db/mock/db_service.go github.com/ZoengYu/order-fast-project/db/sqlc DBService

proto:
	rm -f pb/*.proto
	rm -rf docs/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=order_fast \
    proto/*.proto

evans:
	evans --host localhost --port 8082 -r repl

fmt:
	@echo ">> format code style"
	$(GOFMT) -w $$(find . -path ./vendor -prune -o -name '*.go' -print)

.PHONY: postgres createdb dropdb migrateup migratedown sqlc tests mock proto evans fmt
