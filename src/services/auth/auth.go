package auth

import (
	"context"
	"errors"
	"math/rand"
	"template/src/filter"
	"template/src/models"
	"template/src/repositories/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Interface interface {
	Register(ctx context.Context, input models.Query[models.UserInput]) error
	Login(ctx context.Context, input models.Login) ([]models.User, string, error)
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
	_, count, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.Model.UserName,
		},
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already taken by another user")
	}

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

func (s *authService) Login(ctx context.Context, input models.Login) ([]models.User, string, error) {

	users, _, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.UserName,
		},
	})
	if err != nil {
		return []models.User{}, "", err
	}
	if len(users) == 0 {
		return []models.User{}, "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(input.Password))
	if err != nil {
		return []models.User{}, "", errors.New("wrong password")
	}

	token, err := s.generateToken(int(users[0].Id), users[0].UserName)
	if err != nil {
		return []models.User{}, "", err
	}

	return users, token, nil
}

func (s *authService) generateToken(userId int, userName string) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = userId
	claim["time"] = time.Now().Add(time.Hour * 24 * 3)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(models.GetSecret())

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
