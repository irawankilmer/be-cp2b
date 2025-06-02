package repository

import (
	"be-cp2b/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckEmail(email string) (*domain.User, error)
	UpdateTokenVersion(userID uint) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CheckEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *authRepository) UpdateTokenVersion(userID uint) error {
	return r.db.Model(&domain.User{}).
		Where("id = ?", userID).
		Update("token_version", uuid.New().String()).Error
}
