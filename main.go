package main

import (
	"database/sql"
	"template/src/handler"
	"template/src/repositories"
	"template/src/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/template?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	repo := repositories.Init(repositories.Param{Db: db})

	srv := services.Init(services.Param{Repositories: repo})

	hndlr := handler.Init(handler.InitParam{Service: srv})

	hndlr.Run()

}
