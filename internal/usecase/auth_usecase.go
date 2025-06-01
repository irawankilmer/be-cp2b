package usecase

import (
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
	"be-cp2b/pkg/utils"
	"errors"
)

type AuthUsecase interface {
	Login(req request.AuthRequest) (string, error)
}

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo}
}

func (u *authUsecase) Login(req request.AuthRequest) (string, error) {
	user, err := u.repo.CheckEmail(req.Email)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("Email atau password salah!")
	}

	token, err := utils.GenerateJWT(user.ID)
	return token, err
}
