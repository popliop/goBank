build:
	@go build -o bin/gobank.exe cmd/main.go

run: build
	@./bin/gobank.exe

test:
	"go test -v ./... 