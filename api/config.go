package api

import (
	"net/http"

	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/utils"
)

func getConfig(w http.ResponseWriter, r *http.Request) {
	utils.SendOK(w, controllers.GetConfig())
}
