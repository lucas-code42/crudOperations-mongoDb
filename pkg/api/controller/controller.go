package controller

import (
	"net/http"

	"github.com/lucas-code42/testMongoDb/pkg/api/service"
)

type ControllerImp interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.ServiceImp
}

func NewController(service service.ServiceImp) ControllerImp {
	return &controller{
		service: service,
	}
}
