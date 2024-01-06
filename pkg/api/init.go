package api

import (
	"github.com/lucas-code42/testMongoDb/pkg/api/controller"
	"github.com/lucas-code42/testMongoDb/pkg/api/repository"
	"github.com/lucas-code42/testMongoDb/pkg/api/service"
)

func InitDependencies() controller.ControllerImp {
	// camada de baixo nivel
	repo := repository.NewBookRepository("mongo")

	// camada de domínio
	service := service.NewBookService(repo)

	// camada da api/controller
	controller := controller.NewController(service)
	return controller
}
