package entities

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Email    string
}

func (u *User) TableName() string {
	return "users"
}
