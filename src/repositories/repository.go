package repositories

import (
	"database/sql"
	"template/src/repositories/user"
)

type Repositories struct {
	User user.Interface
}

type Param struct {
	Db *sql.DB
}

func Init(param Param) *Repositories {
	return &Repositories{
		User: user.Init(user.Param{Db: param.Db, TableName: "user"}),
	}
}
