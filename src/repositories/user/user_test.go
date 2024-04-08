package user

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"regexp"
	"template/src/filter"
	"template/src/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yudhaananda/go-common/formatter"
	"github.com/yudhaananda/go-common/paging"
	querybuilder "github.com/yudhaananda/go-common/query_builder"
)

func TestCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO user () VALUES ()")

	type args struct {
		ctx    context.Context
		models models.UserInput
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantErr     bool
	}{
		{
			name: "sql begin failed",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin().WillReturnError(err)
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql exec failed",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectPrepare(query)
				sqlMock.ExpectExec(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit success",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectPrepare(query)
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit()
				return sqlServer, err
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			err = init.Create(tt.args.ctx, tt.args.models, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE user SET  WHERE ")

	type args struct {
		ctx    context.Context
		models models.UserInput
		id     int
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantErr     bool
	}{
		{
			name: "sql begin failed",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin().WillReturnError(err)
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql exec failed",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql no row affected",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(0))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit failed",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit().WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit success",
			args: args{
				ctx:    context.Background(),
				models: models.UserInput{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectPrepare(query)
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit()
				return sqlServer, err
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			err = init.Update(tt.args.ctx, tt.args.models, tt.args.id, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tempModels := models.User{}
	member := querybuilder.BuildTableMember(tempModels)
	query := regexp.QuoteMeta("SELECT " + member + " FROM user WHERE 1=1")
	queryCount := regexp.QuoteMeta("SELECT COUNT(*) FROM user")
	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx    context.Context
		models paging.Paging[filter.UserFilter]
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantUser    []models.User
		wantCount   int
		wantErr     bool
	}{
		{
			name: "sql count query failed",
			args: args{
				ctx:    context.Background(),
				models: paging.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectQuery(queryCount).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantUser: []models.User{},
			wantErr:  true,
		},
		{
			name: "sql query failed",
			args: args{
				ctx:    context.Background(),
				models: paging.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectPrepare(queryCount)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				sqlMock.ExpectPrepare(query)
				sqlMock.ExpectQuery(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr:   true,
			wantUser:  []models.User(nil),
			wantCount: 1,
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				models: paging.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectPrepare(queryCount)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				row := sqlMock.NewRows([]string{"id", "user_name", "password", "name", "birthdate", "age", "status", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"})
				row.AddRow(1, "test", "test", formatter.Null[string]{Valid: true, Data: "test"}, formatter.Null[time.Time]{Valid: true, Data: mockTime}, 5, 1, formatter.Null[time.Time]{Valid: true, Data: mockTime}, 1, formatter.Null[time.Time]{Valid: true, Data: mockTime}, 1, formatter.Null[time.Time]{Valid: true, Data: mockTime}, 1)
				sqlMock.ExpectPrepare(query)
				sqlMock.ExpectQuery(query).WillReturnRows(row)
				return sqlServer, err
			},
			wantUser: []models.User{
				{
					Id:       1,
					UserName: "test",
					Password: "test",
					Name: formatter.Null[string]{
						Valid: true,
						Data:  "test",
					},
					Birthdate: formatter.Null[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					Age: formatter.Null[int64]{
						Data:  5,
						Valid: true,
					},
					Status: 1,
					CreatedAt: formatter.Null[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					UpdatedAt: formatter.Null[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					DeletedAt: formatter.Null[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					CreatedBy: formatter.Null[int64]{
						Data:  1,
						Valid: true,
					},
					UpdatedBy: formatter.Null[int64]{
						Data:  1,
						Valid: true,
					},
					DeletedBy: formatter.Null[int64]{
						Data:  1,
						Valid: true,
					},
				},
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			users, count, err := init.Get(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantUser, users)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}
