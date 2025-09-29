package container

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"

	"github.com/base-go/backend/internal/attendances"
	"github.com/base-go/backend/pkg/cache"
	"github.com/base-go/backend/pkg/router"
	"github.com/base-go/backend/pkg/server"
	"github.com/base-go/backend/pkg/supabase"
)

func New() (*dig.Container, error) {
	// we use go.uber.org/dig for autowire dependencies
	container := dig.New()

	// provide dependencies injection

	if err := container.Provide(cache.NewCache); err != nil {
		return nil, err
	}

	// supabase
	if err := container.Provide(supabase.NewSupabaseClient); err != nil {
		return nil, err
	}

	// attendances
	if err := container.Provide(attendances.NewRepository); err != nil {
		return nil, err
	}

	if err := container.Provide(attendances.NewService); err != nil {
		return nil, err
	}

	if err := container.Provide(attendances.NewHandler); err != nil {
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
