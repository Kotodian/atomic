package service

type User interface {
	// 登录
	Login(string, string) error
	// 创建博客
	CreateBlog(Blog) error
	// 删除博客
	DeleteBlog(Blog) error
	// 收藏博客
	CollectBlog(Blog) error
}
