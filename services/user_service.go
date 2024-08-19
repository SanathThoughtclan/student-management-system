package services

import (
	"context"
	"errors"

	"student-management-system/models"
	"student-management-system/repositories"
	"student-management-system/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Authenticate(ctx context.Context, ID, username, password string) (string, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(ID, username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Register(ctx context.Context, user *models.User) error {
	_, err := s.repo.GetByUsername(ctx, user.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.CreatedOn = time.Now()

	_, err = s.repo.Create(ctx, user)
	return err
}
