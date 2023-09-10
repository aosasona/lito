package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/utils"
)

var server *http.Server

// ServeAdminAPI serves the admin API if enabled
func ServeAdminAPI(port int) error {
	c := chi.NewRouter()

	c.Get("/config", protect(getCurrentConfig))

	server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      c,
	}

	return server.ListenAndServe()
}

func Shutdown(ctx context.Context) error {
	return server.Shutdown(ctx)
}

func getCurrentConfig(w http.ResponseWriter, r *http.Request) {
	utils.SendOK(w, controllers.GetConfig())
}
