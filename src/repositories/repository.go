package repositories

import (
	"database/sql"
	"template/src/repositories/auth"
	"template/src/repositories/user"
)

type Repositories struct {
	Auth auth.Interface
	User user.Interface
}

type Param struct {
	Db *sql.DB
}

func Init(param Param) *Repositories {
	return &Repositories{
		Auth: auth.Init(),
		User: user.Init(user.Param{Db: param.Db, TableName: "user"}),
	}
}
