package ports

type AuthService interface {
	SignUp(user *domain.User) (*domain.User, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}

type AuthRepository interface {
	SignUp(user *domain.User) (*domain.User, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}
