package services

import (
	"auth/internal/core/domain"
	"auth/internal/core/ports"
)

type AuthService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignUp(user *domain.Users) (*domain.Users, error) {
	return a.repo.SignUp(user)
}

func (a *AuthService) SignIn(username, password string) (*repositorys.LoginResponse, error) {
	return a.repo.SignIn(username, password)
}
