package collection

type Collection struct {
	Id     int64 `json:"id" gorm:"id"`
	UserId int64 `json:"user_id" gorm:"user_id"`
	BlogId int64 `json:"blog_id" gorm:"blog_id"`
}
