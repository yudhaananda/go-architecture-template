package user

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"regexp"
	"template/src/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO user () VALUES ()")

	type args struct {
		ctx    context.Context
		models models.Query[models.UserInput]
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
				models: models.Query[models.UserInput]{},
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
				models: models.Query[models.UserInput]{},
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
				models: models.Query[models.UserInput]{},
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
				models: models.Query[models.UserInput]{},
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
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
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
			err = init.Create(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
