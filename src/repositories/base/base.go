package base

import (
	"context"
	"database/sql"
	"reflect"
	"template/src/filter"
	"template/src/models"
)

type BaseInterface[T, M, F comparable] interface {
	Get(ctx context.Context, paging filter.Paging[F]) ([]M, int, error)
	Create(ctx context.Context, input models.Query[T]) error
	Update(ctx context.Context, input models.Query[T], id int) error
}

type BaseRepository[T, M, F comparable] struct {
	Db        *sql.DB
	TableName string
}

func (r *BaseRepository[T, M, F]) Update(ctx context.Context, input models.Query[T], id int) error {
	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}

	updateQuery := input.BuildUpdateQuery(id)

	if _, err = tx.ExecContext(ctx, Update+r.TableName+updateQuery); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *BaseRepository[T, M, F]) Create(ctx context.Context, input models.Query[T]) error {
	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}

	createQuery := input.BuildCreateQuery()

	if _, err = tx.ExecContext(ctx, Create+r.TableName+createQuery); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *BaseRepository[T, M, F]) Get(ctx context.Context, paging filter.Paging[F]) ([]M, int, error) {
	users := []M{}
	var count int

	where := paging.QueryBuilder()
	pagination := paging.PaginationQuery()

	rowCount, err := r.Db.QueryContext(ctx, Count+r.TableName+where)
	if err != nil {
		return users, 0, err
	}

	defer rowCount.Close()
	for rowCount.Next() {
		err = rowCount.Scan(&count)
		if err != nil {
			return users, count, err
		}
	}
	row, err := r.Db.QueryContext(ctx, Select+r.TableName+where+pagination)
	if err != nil {
		return users, count, err
	}

	defer row.Close()
	for row.Next() {
		var user M

		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := row.Scan(columns...)
		if err != nil {
			return users, count, err
		}
		users = append(users, user)
	}

	return users, count, nil
}
