.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./main.go -o ./docs/swagger --parseInternal

.PHONY: initialize
initialize:
	@go get "github.com/gin-gonic/gin"
	@go get "github.com/gin-contrib/cors"
	@go get "github.com/swaggo/files"
	@go get "github.com/alecthomas/template"
	@go get "github.com/swaggo/gin-swagger"
	@go get "github.com/dgrijalva/jwt-go"

.PHONY: run
run: swaggo
	@go run main.go
