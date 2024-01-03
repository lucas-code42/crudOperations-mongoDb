package router

import (
	"net/http"

	"github.com/lucas-code42/testMongoDb/pkg/api/controller"
)

func StartRouters(controller controller.ControllerImp) {
	http.HandleFunc("/", controller.Get)
	http.ListenAndServe(":8080", nil)
}
