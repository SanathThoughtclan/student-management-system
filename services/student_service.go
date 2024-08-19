package services

import (
	"context"
	"student-management-system/models"
	"student-management-system/repositories"
	"time"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) Create(ctx context.Context, student *models.Student, username string) error {
	student.CreatedBy = username
	student.CreatedOn = time.Now()
	student.UpdatedBy = ""
	student.UpdatedOn = time.Now()
	return s.repo.Create(ctx, student)
}

func (s *StudentService) GetAll(ctx context.Context) ([]*models.Student, error) {
	return s.repo.GetAll(ctx)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (*models.Student, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *StudentService) Update(ctx context.Context, id string, student *models.Student) error {
	// Just pass the student to the repository
	return s.repo.Update(ctx, id, student)
}

func (s *StudentService) Delete(ctx context.Context, idStr string) error {
	return s.repo.Delete(ctx, idStr)
}
