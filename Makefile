test:
	go test -v ./cmd/... ./internal/... --race

build:
	go clean
	go build -o counting-request-server ./cmd/
