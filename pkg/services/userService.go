package services

import (
	"github.com/JcsnP/gin-clean/pkg/models"
	"github.com/JcsnP/gin-clean/pkg/repositories"
)

type UserService struct {
	userRepo		*repositories.UserRepository
}

func NewUserRepository(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAll() ([]models.User, int64, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
} 

func (s *UserService) Create(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) Delete(user *models.User, id uint) error {
	return s.userRepo.Delete(user, id)
}

func (s *UserService) Update(user *models.User) error {
	return s.userRepo.Update(user)
}