package repositories

import (
	"github.com/JcsnP/gin-clean/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db			*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, int64, error) {
	var users []models.User
	var count int64
	r.db.Find(&users).Count(&count)
	
	return users, count, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil 
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Delete(user *models.User, id uint) error {
	return r.db.Delete(&user, id).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}