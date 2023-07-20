package user

import (
	"context"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/user"
	"time"
)

type Interface interface {
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, input models.Query[models.UserInput], id int) error
	Create(ctx context.Context, input models.Query[models.UserInput]) error
	Get(ctx context.Context, paging filter.Paging[filter.UserFilter]) ([]models.User, int, error)
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
	input := models.Query[models.UserInput]{
		Model: models.UserInput{
			Status:    -1,
			DeletedAt: Now(),
			DeletedBy: ctx.Value(models.UserKey).(models.User).Id,
		},
	}

	return s.userRepository.Update(ctx, input, id)
}

func (s *userService) Update(ctx context.Context, input models.Query[models.UserInput], id int) error {
	input.Model.UpdatedAt = Now()
	input.Model.UpdatedBy = ctx.Value(string(models.UserKey)).(models.User).Id

	return s.userRepository.Update(ctx, input, id)
}

func (s *userService) Create(ctx context.Context, input models.Query[models.UserInput]) error {
	input.Model.CreatedAt = Now()
	input.Model.CreatedBy = ctx.Value(models.UserKey).(models.User).Id

	return s.userRepository.Create(ctx, input)
}

func (s *userService) Get(ctx context.Context, paging filter.Paging[filter.UserFilter]) ([]models.User, int, error) {
	paging.IsActive = true
	return s.userRepository.Get(ctx, paging)
}
