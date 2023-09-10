package api

import (
	"net/http"

	"go.trulyao.dev/lito/pkg/errors"
	"go.trulyao.dev/lito/pkg/utils"
)

func protect(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle auth here
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			utils.SendErrorResponse(w, errors.NewAPIError("Missing API key", http.StatusUnauthorized))
			return
		}

		handler(w, r)
	}
}
