.PHONY:
.SILENT:

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/main ./cmd/passport/main.go

run: build
	docker-compose up --remove-orphans

test:
	go test -v ./...

swag:
	swag init -g cmd/passport/main.go
