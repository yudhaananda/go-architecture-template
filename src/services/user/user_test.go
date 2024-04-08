package user_test

import (
	"context"
	"template/src/filter"
	"template/src/models"
	mock_user "template/src/repositories/mock/user"
	"template/src/services/user"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yudhaananda/go-common/formatter"
	"github.com/yudhaananda/go-common/paging"
)

func Test_userService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: formatter.NewNull[int64](1)})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Input models.UserInput
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	user.Now = func() time.Time {
		return mockTime
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
			name: "create user error",
			args: args{
				Input: models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Create(context, models.UserInput{
					CreatedBy: context.Value(models.UserKey).(models.User).Id.Data,
					CreatedAt: mockTime,
				}, nil).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "create user success",
			args: args{
				models.UserInput{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Create(context, models.UserInput{
					CreatedBy: context.Value(models.UserKey).(models.User).Id.Data,
					CreatedAt: mockTime,
				}, nil).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Create(context, tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: formatter.NewNull[int64](1)})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Input models.UserInput
		Id    int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	user.Now = func() time.Time {
		return mockTime
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
			name: "update user error",
			args: args{
				Input: models.UserInput{},
				Id:    1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.UserInput{
					UpdatedBy: context.Value(models.UserKey).(models.User).Id.Data,
					UpdatedAt: mockTime,
				}, 1, nil).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "update user success",
			args: args{
				Input: models.UserInput{},
				Id:    1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.UserInput{
					UpdatedBy: context.Value(models.UserKey).(models.User).Id.Data,
					UpdatedAt: mockTime,
				}, 1, nil).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Update(context, tt.args.Input, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: formatter.NewNull[int64](1)})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Id int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	user.Now = func() time.Time {
		return mockTime
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
			name: "delete user error",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.UserInput{
					DeletedBy: context.Value(models.UserKey).(models.User).Id.Data,
					DeletedAt: mockTime,
					Status:    -1,
				}, 1, nil).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "delete user success",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.UserInput{
					DeletedBy: context.Value(models.UserKey).(models.User).Id.Data,
					DeletedAt: mockTime,
					Status:    -1,
				}, 1, nil).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Delete(context, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.Background()

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Paging paging.Paging[filter.UserFilter]
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name      string
		args      args
		mockfunc  func(a args, mock mockfields)
		want      []models.UserDto
		wantCount int
		wantErr   bool
	}{
		{
			name: "get user error",
			args: args{
				paging.Paging[filter.UserFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, paging.Paging[filter.UserFilter]{IsActive: true}).Return([]models.User{}, 0, assert.AnError)
			},
			want:      nil,
			wantCount: 0,
			wantErr:   true,
		},
		{
			name: "get user success",
			args: args{
				paging.Paging[filter.UserFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, paging.Paging[filter.UserFilter]{IsActive: true}).Return([]models.User{
					{},
					{},
				}, 2, nil)
			},
			want: []models.UserDto{
				{},
				{},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			users, count, err := service.Get(context, tt.args.Paging)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, users)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}
