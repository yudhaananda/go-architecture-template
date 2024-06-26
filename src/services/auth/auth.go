package auth

import (
	"context"
	"errors"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/auth"
	"template/src/repositories/user"

	"github.com/yudhaananda/go-common/paging"
)

type Interface interface {
	Register(ctx context.Context, input models.UserInput) error
	Login(ctx context.Context, input models.Login) (*models.UserDto, string, error)
}

type authService struct {
	authRepository auth.Interface
	userRepository user.Interface
}

type Param struct {
	AuthRepository auth.Interface
	UserRepository user.Interface
}

func Init(param Param) *authService {
	return &authService{userRepository: param.UserRepository, authRepository: param.AuthRepository}
}

func (s *authService) Register(ctx context.Context, input models.UserInput) error {
	_, count, err := s.userRepository.Get(ctx, paging.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.UserName,
		},
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already taken by another user")
	}

	password, err := s.authRepository.HashPassword([]byte(input.Password))
	if err != nil {
		return err
	}
	input.Password = password

	err = s.userRepository.Create(ctx, input, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(ctx context.Context, input models.Login) (*models.UserDto, string, error) {

	users, _, err := s.userRepository.Get(ctx, paging.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.UserName,
		},
	})
	if err != nil {
		return nil, "", err
	}
	if len(users) == 0 {
		return nil, "", errors.New("user not found")
	}

	err = s.authRepository.ComparePassword([]byte(users[0].Password.Data), []byte(input.Password))
	if err != nil {
		return nil, "", errors.New("wrong password")
	}

	token, err := s.authRepository.GenerateToken(int(users[0].Id.Data), users[0].UserName.Data)
	if err != nil {
		return nil, "", err
	}

	return users[0].ToDto(), token, nil
}
