.PHONY: test
test:
	@go test --tags unit -parallel 20 -failfast \
		`go list ./... | grep -v mocks | grep -v docs` \
		-race -short -coverprofile=./cov.out
		
.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --parseDependency --parseInternal

.PHONY: initialize
initialize:
	@go install github.com/golang/mock/mockgen@v1.6.0
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
	@go mod tidy
	@./build/app

.PHONY: mock
mock:
	@`go env GOPATH`/bin/mockgen -source src/repositories/$(repositories)/$(repositories).go -destination src/repositories/mock/$(repositories)/$(repositories).go

.PHONY: mock-all
mock-all:
	@make mock repositories=auth
	@make mock repositories=user
