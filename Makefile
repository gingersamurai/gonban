PROJECT=gonban
BINDIR=./bin

SQLITE_DB_PATH=db/taskDB.sqlite
SQLITE_MIGRATIONS_PATH=db/migrations/sqlite/00.sql

POSTGRES_MIGRATIONS_PATH=db/migrations/postgres/00.sql


run:
	go run ${PROJECT}/cmd/server_rest

build:
	go build -o ${BINDIR}/gonban_server_rest ${PROJECT}/cmd/server_rest

test:
	go test ./...



postgres_migrate:
	 goose -dir ./db/migrations up

postgres_init:
	docker run --rm --name some-postgres -p 5432:5432 --env-file .env -d

clean:
	rm -rf ${SQLITE_DB_PATH}
	docker stop some-postgres
