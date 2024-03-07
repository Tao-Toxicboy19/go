package repo

import (
	"auth/internal/core/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *DB) SignUp(user *domain.Users) (*domain.Users, error) {
	req := a.db.Where("username = ? ", user.Username).First(&user)

	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	user.ID = uuid.New().String()
	user.Password = string(hash)

	req = a.db.Create(&user)

	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}

	return user, nil
}
