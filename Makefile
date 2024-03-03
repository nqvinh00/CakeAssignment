build:
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -s" -o ../build/app ./cmd/*.go

lint:
	golangci-lint run -v ./...

.PHONY: build
