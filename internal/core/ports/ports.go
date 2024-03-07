package ports

import (
	"auth/internal/core/domain"
)

type AuthService interface {
	SignUp(user *domain.Users) (*domain.Users, error)
}

type AuthRepository interface {
	SignUp(user *domain.Users) (*domain.Users, error)
}
