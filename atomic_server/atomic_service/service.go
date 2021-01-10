package atomic_service

import (
	"atomic/internal/cache"
	"gorm.io/gorm"
)

type BlogService struct {
	DB    *gorm.DB     `inject:""`
	Cache *cache.Cache `inject:""`
	// FIXME 自行扩展
}

type UserService struct {
	DB    *gorm.DB     `inject:""`
	Cache *cache.Cache `inject:""`
	// FIXME 自行扩展
}
