test:
	go test -v ./cmd/... ./internal/... --race

build:
	go build -o counting-request-server ./cmd/
