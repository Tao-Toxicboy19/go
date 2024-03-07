package ports

import "auth/internal/core/domain"

type AuthService interface {
	SignUp(user *domain.Users) (*domain.Users, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}

type AuthRepository interface {
	SignUp(user *domain.Users) (*domain.Users, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}
