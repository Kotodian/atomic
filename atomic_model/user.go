package atomic_model

type User struct {
	Id       int64  `gorm:"id"`
	Username string `gorm:"username"`
	Nickname string `gorm:"nickname"`
	Email    string `gorm:"email"`
	Phone    string `gorm:"phone"`
}
