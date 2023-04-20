PROJECT=github.com/gingersamurai/gonban
BINDIR=./bin

SQLITE_DB_PATH=db/taskDB.sqlite
SQLITE_MIGRATIONS_PATH=db/migrations/sqlite/00.sql

POSTGRES_MIGRATIONS_PATH=db/migrations/postgres/00.sql


run: build
	go run ${PROJECT}/cmd/server_rest

build:
	go build -o ${BINDIR}/gonban_server_rest ${PROJECT}/cmd/server_rest

test:
	go test ./...


sqlite_migrate:
	sqlite3  ${SQLITE_DB_PATH} '.read ${SQLITE_MIGRATIONS_PATH}'

sqlite_init:
	touch ${SQLITE_DB_PATH}


postgres_migrate:
	psql "host=localhost user=postgres password=15092003" -f ${POSTGRES_MIGRATIONS_PATH}

postgres_init:
	docker run --rm --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=15092003 -d postgres


clean:
	rm -rf ${SQLITE_DB_PATH}
	docker stop some-postgres
