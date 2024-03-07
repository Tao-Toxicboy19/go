package domain

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
