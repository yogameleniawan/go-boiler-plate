package attendances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/base-go/backend/internal/shared/models"
)

type Service interface {
	Create(ctx context.Context, attendance models.Attendance) (*models.Attendance, int, error)
	GetAll(ctx context.Context) ([]models.Attendance, int, error)
	GetByID(ctx context.Context, id int64) (*models.Attendance, int, error)
	Update(ctx context.Context, id int64, attendance models.Attendance) (*models.Attendance, int, error)
	Delete(ctx context.Context, id int64) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, attendance models.Attendance) (*models.Attendance, int, error) {
	data, err := s.repo.Create(ctx, attendance)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return data, http.StatusCreated, nil
}

func (s *service) GetAll(ctx context.Context) ([]models.Attendance, int, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return data, http.StatusOK, nil
}

func (s *service) GetByID(ctx context.Context, id int64) (*models.Attendance, int, error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if data == nil {
		return nil, http.StatusNotFound, fmt.Errorf("attendance with id %d not found", id)
	}
	return data, http.StatusOK, nil
}

func (s *service) Update(ctx context.Context, id int64, attendance models.Attendance) (*models.Attendance, int, error) {
	data, err := s.repo.Update(ctx, id, attendance)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if data == nil {
		return nil, http.StatusNotFound, fmt.Errorf("attendance with id %d not found for update", id)
	}
	return data, http.StatusOK, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int, error) {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
