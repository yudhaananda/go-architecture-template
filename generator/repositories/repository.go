package repositories

import (
	"codegen/repositories/generate"
	"codegen/repositories/zipping"
)

type Repositories struct {
	Zipping  zipping.Interface
	Generate generate.Interface
}

type Param struct {
}

func Init(param Param) *Repositories {
	return &Repositories{
		Zipping:  zipping.Init(),
		Generate: generate.Init(),
	}
}
