package services

import (
	"errors"

	"github.com/ferryku8/project-management/models"
	"github.com/ferryku8/project-management/repositories"
	"github.com/ferryku8/project-management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserServices(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}

	hased, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid Credential")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid Credential")
	}
	return user, nil
}
