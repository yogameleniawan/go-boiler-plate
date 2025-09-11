package health

import (
	"context"

	"github.com/absendulu-project/backend/pkg/cache"
)

type Service interface {
	Health(ctx context.Context) HealthResponse
}

type service struct {
	Repository Repository
	cache      cache.Cache
}

func NewService(cache cache.Cache, repository Repository) Service {
	return &service{
		Repository: repository,
		cache:      cache,
	}
}

func (s *service) Health(ctx context.Context) HealthResponse {
	resp := HealthResponse{
		Status: "Service is healthy",
	}
	data := HealthData{}

	data.Postgres = s.Repository.Ping(ctx)

	if !data.Postgres {
		resp.Status = "Service is unhealthy"
	}

	if err := s.cache.Ping(ctx); err != nil {
		resp.Status = "Service is unhealthy"
	} else {
		data.Cache = true
	}

	resp.Dependencies = data
	return resp
}
