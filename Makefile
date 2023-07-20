.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --parseInternal

.PHONY: initialize
initialize:
	@go get "github.com/go-sql-driver/mysql"
	@go get "github.com/gin-gonic/gin"
	@go get "github.com/gin-contrib/cors"
	@go get "github.com/swaggo/files"
	@go get "github.com/alecthomas/template"
	@go get "github.com/swaggo/gin-swagger"
	@go get "github.com/dgrijalva/jwt-go"
	@go get "github.com/joho/godotenv"
	@go get "github.com/DATA-DOG/go-sqlmock"

.PHONY: build
build:
	@go build -o ./build/app ./src/cmd

.PHONY: run
run: swaggo build
	@./build/app
