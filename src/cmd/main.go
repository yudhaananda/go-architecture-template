package main

import (
	"database/sql"
	"fmt"
	"log"

	"template/src/handler"
	"template/src/middleware"
	"template/src/repositories"
	"template/src/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yudhaananda/go-common/env"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	env, err := env.SetEnv()
	if err != nil {
		log.Fatal(err)
		return
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", env.DB_USER, env.DB_PASS, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	db, err := sql.Open(env.DB_TYPE, dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	repo := repositories.Init(repositories.Param{Db: db})

	srv := services.Init(services.Param{Repositories: repo})

	midlwre := middleware.Init(middleware.InitParam{Service: srv})

	hndlr := handler.Init(handler.InitParam{Service: srv, Middleware: midlwre})

	hndlr.Run()

}
