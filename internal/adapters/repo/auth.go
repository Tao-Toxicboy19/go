package repo

func (a *DB) SignUp(user *domain.User) (*domain.User, error) {
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
