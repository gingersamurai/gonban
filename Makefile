PROJECT=github.com/gingersamurai/gonban
BINDIR=./bin

run: build
	go run ${PROJECT}/cmd/server_rest

build:
	go build -o ${BINDIR}/gonban_server_rest ${PROJECT}/cmd/server_rest

test:
	go test ./...