package repository

import (
	"be-cp2b/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckEmail(email string) (*domain.User, error)
	IncrementTokenVersion(user *domain.User) error
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

func (r *authRepository) IncrementTokenVersion(user *domain.User) error {
	newTokenVersion := uuid.New().String()
	user.TokenVersion = newTokenVersion
	return r.db.Save(user).Error
}
