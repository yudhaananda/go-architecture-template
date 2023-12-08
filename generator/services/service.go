package services

import (
	"codegen/repositories"
	"codegen/services/generate"
)

type Services struct {
	Generate generate.Interface
}

type Param struct {
	Repositories *repositories.Repositories
}

func Init(param Param) *Services {
	return &Services{
		Generate: generate.Init(generate.Params{
			GenerateRepo: param.Repositories.Generate,
			Zipping:      param.Repositories.Zipping,
		}),
	}
}
