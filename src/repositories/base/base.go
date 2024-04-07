package base

import (
	"context"
	"fmt"

	"github.com/yudhaananda/go-common/db/sql"
	"github.com/yudhaananda/go-common/paging"
	querybuilder "github.com/yudhaananda/go-common/query_builder"
)

type BaseInterface[T, M, F comparable] interface {
	Get(ctx context.Context, paging paging.Paging[F]) ([]M, int, error)
	Create(ctx context.Context, input T, trx *sql.Tx) error
	Update(ctx context.Context, input T, id int, trx *sql.Tx) error
}

type BaseRepository[T, M, F comparable] struct {
	Db        *sql.DBSql[M]
	TableName string
}

func (r *BaseRepository[T, M, F]) Update(ctx context.Context, input T, id int, trx *sql.Tx) error {
	updateQuery, args := querybuilder.BuildUpdateQuery(id, input)

	if err := r.Db.ExecContext(ctx, Update+r.TableName+updateQuery, trx, args...); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository[T, M, F]) Create(ctx context.Context, input T, trx *sql.Tx) error {
	createQuery, args := querybuilder.BuildCreateQuery(input)

	if err := r.Db.ExecContext(ctx, Create+r.TableName+createQuery, trx, args...); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository[T, M, F]) Get(ctx context.Context, paging paging.Paging[F]) ([]M, int, error) {

	var (
		where, countArgs = paging.QueryBuilder()
		pagination, args = paging.PaginationQuery(countArgs)
		tempModels       M
		member           = querybuilder.BuildTableMember(tempModels)
		query            = fmt.Sprintf(Select, member)
		models           = []M{}
		count            int
	)

	err := r.Db.CountContext(ctx, &count, Count+r.TableName+where, countArgs...)
	if err != nil {
		return models, count, err
	}

	models, err = r.Db.GetContext(ctx, query+r.TableName+where+pagination, args...)
	if err != nil {
		return models, count, err
	}

	return models, count, nil
}
