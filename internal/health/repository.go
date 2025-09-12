package health

import (
	"context"
)

type Repository interface {
	Ping(ctx context.Context) bool
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Ping(ctx context.Context) bool {

	return true
}
