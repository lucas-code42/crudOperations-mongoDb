package router

import (
	"net/http"

	"github.com/lucas-code42/testMongoDb/pkg/api/controller"
)

var routerController controller.ControllerImp

func StartRouters(controller controller.ControllerImp) {
	routerController = controller

	http.HandleFunc("/", mainRouter)
	http.ListenAndServe(":8080", nil)
}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		routerController.Create(w, r)
		return
	}
}
