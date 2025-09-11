package health

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Ping(ctx context.Context) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Ping(ctx context.Context) bool {

	db, err := r.db.DB()
	if err != nil {
		return false
	}

	if err := db.Ping(); err != nil {
		return false
	}

	return true
}
