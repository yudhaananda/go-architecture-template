package auth

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yudhaananda/go-common/env"
	"golang.org/x/crypto/bcrypt"
)

type Interface interface {
	HashPassword(pwd []byte) (string, error)
	ComparePassword(hashedPassword, inputPassword []byte) error
	GenerateToken(userId int, userName string) (string, error)
}

type authRepository struct {
}

func Init() Interface {
	return &authRepository{}
}

func (r *authRepository) HashPassword(pwd []byte) (string, error) {
	key := rand.Intn(9)
	password, err := bcrypt.GenerateFromPassword(pwd, key)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (r *authRepository) ComparePassword(hashedPassword, inputPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, inputPassword)
}

func (s *authRepository) GenerateToken(userId int, userName string) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = userId
	claim["time"] = time.Now().Add(time.Hour * 24 * 3)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secret, err := env.GetSecret()
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(secret)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
