package service

import (
	"github.com/5MofDream/apollo/app/models"
	"github.com/5MofDream/apollo/app/repository"
)

type ExampleService struct {
	Base
}

func (es *ExampleService) DoExample() []models.Users {
	re := new(repository.Example)
	re.DI(es.Container)
	return re.GetList()
}
