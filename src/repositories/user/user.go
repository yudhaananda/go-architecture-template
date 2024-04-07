package user

import (
	"database/sql"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/base"

	sql_db "github.com/yudhaananda/go-common/db/sql"
)

type Interface interface {
	base.BaseInterface[models.UserInput, models.User, filter.UserFilter]
}

type userRepository struct {
	base.BaseRepository[models.UserInput, models.User, filter.UserFilter]
}
type Param struct {
	Db        *sql.DB
	TableName string
}

func Init(param Param) Interface {
	return &userRepository{
		BaseRepository: base.BaseRepository[models.UserInput, models.User, filter.UserFilter]{
			Db: &sql_db.DBSql[models.User]{
				DB: param.Db,
			},
			TableName: param.TableName,
		},
	}
}
