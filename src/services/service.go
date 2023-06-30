package services

import (
	"template/src/repositories"
	"template/src/services/auth"
	"template/src/services/jwt"
	"template/src/services/user"
)

type Services struct {
	User user.Interface
	Jwt  jwt.Interface
	Auth auth.Interface
}

type Param struct {
	Repositories *repositories.Repositories
}

func Init(param Param) *Services {
	return &Services{
		User: user.Init(user.Param{UserRepository: param.Repositories.User}),
		Jwt:  jwt.Init(),
		Auth: auth.Init(auth.Param{UserRepository: param.Repositories.User}),
	}
}
