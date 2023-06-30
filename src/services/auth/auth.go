package auth

import (
	"context"
	"errors"
	"math/rand"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/user"

	"golang.org/x/crypto/bcrypt"
)

type Interface interface {
	Register(ctx context.Context, input models.Query[models.UserInput]) error
	Login(ctx context.Context, input models.Login) (models.User, error)
}

type authService struct {
	userRepository user.Interface
}

type Param struct {
	UserRepository user.Interface
}

func Init(param Param) *authService {
	return &authService{userRepository: param.UserRepository}
}

func (s *authService) Register(ctx context.Context, input models.Query[models.UserInput]) error {

	key := rand.Intn(9)
	password, err := bcrypt.GenerateFromPassword([]byte(input.Model.Password), key)
	if err != nil {
		return err
	}
	input.Model.Password = string(password)

	err = s.userRepository.Create(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(ctx context.Context, input models.Login) (models.User, error) {

	users, _, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.UserName,
		},
	})
	if err != nil {
		return models.User{}, err
	}
	if len(users) == 0 {
		return models.User{}, errors.New("user not found")
	}

	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {

		return models.User{}, err
	}
	return user, nil
}
