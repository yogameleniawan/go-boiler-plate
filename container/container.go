package container

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"

	"github.com/absendulu-project/backend/internal/health"
	"github.com/absendulu-project/backend/pkg/cache"
	"github.com/absendulu-project/backend/pkg/router"
	"github.com/absendulu-project/backend/pkg/server"
)

func New() (*dig.Container, error) {
	// we use go.uber.org/dig for autowire dependencies
	container := dig.New()

	// provide dependencies injection

	if err := container.Provide(cache.NewCache); err != nil {
		return nil, err
	}

	// health
	if err := container.Provide(health.NewRepository); err != nil {
		return nil, err
	}

	if err := container.Provide(health.NewService); err != nil {
		return nil, err
	}

	if err := container.Provide(health.NewHandler); err != nil {
		return nil, err
	}

	// other domain

	// end

	if err := container.Provide(router.SetupRoutes); err != nil {
		return nil, err
	}

	if err := container.Provide(ProvideHttpServer); err != nil {
		return nil, err
	}

	return container, nil
}

func ProvideHttpServer(mux *chi.Mux) (server.Server, error) {
	svr := server.New()
	svr.WithRoute(mux)
	return svr, nil
}
