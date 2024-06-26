package user

import (
	"context"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/user"
	"time"

	"github.com/yudhaananda/go-common/paging"
)

type Interface interface {
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, input models.UserInput, id int) error
	Create(ctx context.Context, input models.UserInput) error
	Get(ctx context.Context, paging paging.Paging[filter.UserFilter]) ([]models.UserDto, int, error)
}

type userService struct {
	userRepository user.Interface
}

type Param struct {
	UserRepository user.Interface
}

func Init(param Param) Interface {
	return &userService{
		userRepository: param.UserRepository,
	}
}

var Now = time.Now

func (s *userService) Delete(ctx context.Context, id int) error {
	input := models.UserInput{
		Status:    -1,
		DeletedAt: Now(),
		DeletedBy: ctx.Value(models.UserKey).(models.User).Id.Data,
	}

	return s.userRepository.Update(ctx, input, id, nil)
}

func (s *userService) Update(ctx context.Context, input models.UserInput, id int) error {
	input.UpdatedAt = Now()
	input.UpdatedBy = ctx.Value(string(models.UserKey)).(models.User).Id.Data

	return s.userRepository.Update(ctx, input, id, nil)
}

func (s *userService) Create(ctx context.Context, input models.UserInput) error {
	input.CreatedAt = Now()
	input.CreatedBy = ctx.Value(models.UserKey).(models.User).Id.Data

	return s.userRepository.Create(ctx, input, nil)
}

func (s *userService) Get(ctx context.Context, paging paging.Paging[filter.UserFilter]) ([]models.UserDto, int, error) {
	paging.IsActive = true

	users, count, err := s.userRepository.Get(ctx, paging)
	if err != nil {
		return nil, 0, err
	}

	res := make([]models.UserDto, 0)
	for _, user := range users {
		res = append(res, *user.ToDto())
	}

	return res, count, nil
}
