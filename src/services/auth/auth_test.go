package auth_test

import (
	"context"
	"template/src/filter"
	"template/src/models"
	mock_auth "template/src/repositories/mock/auth"
	mock_user "template/src/repositories/mock/user"
	"template/src/services/auth"
	"template/src/services/user"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yudhaananda/go-common/formatter"
	"github.com/yudhaananda/go-common/paging"
)

func Test_authService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_user.NewMockInterface(ctrl)
	authRepo := mock_auth.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
		auth *mock_auth.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
		auth: authRepo,
	}
	params := auth.Param{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
	service := auth.Init(params)
	type args struct {
		Input models.UserInput
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "get user error",
			args: args{
				Input: models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), gomock.Any()).Return([]models.User{}, 0, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "get user",
			args: args{
				Input: models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), gomock.Any()).Return([]models.User{
					{},
				}, 1, nil)
			},
			wantErr: true,
		},
		{
			name: "hash password error",
			args: args{
				Input: models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page:   1,
					Take:   1,
					Filter: filter.UserFilter{},
				}).Return([]models.User{}, 0, nil)
				mock.auth.EXPECT().HashPassword([]byte("")).Return("", assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "create user error",
			args: args{
				Input: models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page:   1,
					Take:   1,
					Filter: filter.UserFilter{},
				}).Return([]models.User{}, 0, nil)
				mock.auth.EXPECT().HashPassword([]byte("")).Return("password", nil)
				mock.user.EXPECT().Create(context.Background(), models.UserInput{
					Password: "password",
				}, nil).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "register user success",
			args: args{
				models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page: 1,
					Take: 1,
					Filter: filter.UserFilter{
						UserName: "",
					},
				}).Return([]models.User{}, 0, nil)
				mock.auth.EXPECT().HashPassword([]byte("")).Return("password", nil)
				mock.user.EXPECT().Create(context.Background(), models.UserInput{
					Password: "password",
				}, nil).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Register(context.Background(), tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("auth.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_authService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_user.NewMockInterface(ctrl)
	authRepo := mock_auth.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
		auth *mock_auth.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
		auth: authRepo,
	}
	params := auth.Param{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
	service := auth.Init(params)
	type args struct {
		Input models.Login
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name      string
		args      args
		mockfunc  func(a args, mock mockfields)
		wantErr   bool
		wantUser  *models.UserDto
		wantToken string
	}{
		{
			name: "get user error",
			args: args{
				Input: models.Login{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), gomock.Any()).Return([]models.User{}, 0, assert.AnError)
			},
			wantUser: nil,
			wantErr:  true,
		},
		{
			name: "doesnt get user",
			args: args{
				Input: models.Login{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), gomock.Any()).Return([]models.User{}, 0, nil)
			},
			wantUser: nil,
			wantErr:  true,
		},
		{
			name: "compare password error",
			args: args{
				Input: models.Login{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page:   1,
					Take:   1,
					Filter: filter.UserFilter{},
				}).Return([]models.User{
					{},
				}, 1, nil)
				mock.auth.EXPECT().ComparePassword([]byte(""), []byte("")).Return(assert.AnError)
			},
			wantUser: nil,
			wantErr:  true,
		},
		{
			name: "generate token error",
			args: args{
				Input: models.Login{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page:   1,
					Take:   1,
					Filter: filter.UserFilter{},
				}).Return([]models.User{
					{
						Id:       formatter.NewNull[int64](1),
						UserName: formatter.NewNull[string]("test"),
					},
				}, 1, nil)
				mock.auth.EXPECT().ComparePassword([]byte(""), []byte("")).Return(nil)
				mock.auth.EXPECT().GenerateToken(1, "test").Return("", assert.AnError)
			},
			wantUser: nil,
			wantErr:  true,
		},
		{
			name: "login success",
			args: args{
				Input: models.Login{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context.Background(), paging.Paging[filter.UserFilter]{
					Page:   1,
					Take:   1,
					Filter: filter.UserFilter{},
				}).Return([]models.User{
					{
						Id:       formatter.NewNull[int64](1),
						UserName: formatter.NewNull[string]("test"),
					},
				}, 1, nil)
				mock.auth.EXPECT().ComparePassword([]byte(""), []byte("")).Return(nil)
				mock.auth.EXPECT().GenerateToken(1, "test").Return("token", nil)
			},
			wantUser: &models.UserDto{
				Id:       1,
				UserName: "test",
			},
			wantToken: "token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			user, token, err := service.Login(context.Background(), tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("auth.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, user, tt.wantUser)
			assert.Equal(t, token, tt.wantToken)
		})
	}
}
