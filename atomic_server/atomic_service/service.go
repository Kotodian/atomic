package atomic_service

import "gorm.io/gorm"

type BlogService struct {
	DB *gorm.DB `inject:""`
	// FIXME 自行扩展
}

type UserService struct {
	DB *gorm.DB `inject:""`
	// FIXME 自行扩展
}
