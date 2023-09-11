package api

import (
	"net/http"

	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/errors"
	"go.trulyao.dev/lito/pkg/utils"
)

func protect(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if !controllers.ValidateAPIKey(apiKey) {
			utils.SendErrorResponse(w, errors.NewAPIError("Invalid API key", http.StatusUnauthorized))
			return
		}

		handler(w, r)
		return
	}
}
