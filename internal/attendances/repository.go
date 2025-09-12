package attendances

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/absendulu-project/backend/internal/shared/models"
	"github.com/supabase-community/postgrest-go"
	supabase "github.com/supabase-community/supabase-go"
)

type Repository interface {
	Create(ctx context.Context, attendance models.Attendance) (*models.Attendance, error)
	GetAll(ctx context.Context) ([]models.Attendance, error)
	GetByID(ctx context.Context, id int64) (*models.Attendance, error)
	Update(ctx context.Context, id int64, attendance models.Attendance) (*models.Attendance, error)
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	db *supabase.Client
}

func NewRepository(db *supabase.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, attendance models.Attendance) (*models.Attendance, error) {
	var result []models.Attendance

	data, _, err := r.db.From("attendances").Insert(attendance, false, "", "", "").Execute()
	if err != nil {
		log.Printf("Error creating attendance: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Printf("Error unmarshaling created attendance data: %v", err)
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, fmt.Errorf("did not get back created record")
}

func (r *repository) GetAll(ctx context.Context) ([]models.Attendance, error) {
	var results []models.Attendance

	data, _, err := r.db.From("attendances").
		Select("*", "0", false).
		Order("created_at", &postgrest.OrderOpts{
			Ascending: false,
		}).
		Execute()

	if err != nil {
		log.Printf("Error fetching all attendances: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(data, &results); err != nil {
		log.Printf("Error unmarshaling attendances data: %v", err)
		return nil, err
	}

	return results, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (*models.Attendance, error) {
	var result []models.Attendance

	data, _, err := r.db.From("attendances").
		Select("*", "0", false).
		Eq("id", fmt.Sprintf("%d", id)).
		Execute()
	if err != nil {
		log.Printf("Error fetching attendance by ID: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Printf("Error unmarshaling attendance data by ID: %v", err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *repository) Update(ctx context.Context, id int64, attendance models.Attendance) (*models.Attendance, error) {
	var result []models.Attendance

	data, _, err := r.db.From("attendances").
		Update(attendance, "", "").
		Eq("id", fmt.Sprintf("%d", id)).
		Execute()
	if err != nil {
		log.Printf("Error updating attendance: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Printf("Error unmarshaling updated attendance data: %v", err)
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, fmt.Errorf("did not get back updated record")
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	_, _, err := r.db.From("attendances").
		Delete("", "").
		Eq("id", fmt.Sprintf("%d", id)).
		Execute()

	if err != nil {
		log.Printf("Error deleting attendance: %v", err)
		return err
	}

	return nil
}
