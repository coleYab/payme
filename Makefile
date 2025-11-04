run:
	@swag init
	@go run main.go

build:
	@go build -o payme main.go
