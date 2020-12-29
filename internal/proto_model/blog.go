package proto_model

import (
	"atomic/atomic_model/blog"
	proto "atomic/atomic_proto/blog"
)

func CommonBlog(pbBlog *proto.CommonBlog) *blog.CommonBlog {
	return &blog.CommonBlog{
		Id:      pbBlog.Id,
		UserId:  pbBlog.UserId,
		Title:   pbBlog.Title,
		Content: protoBlogContentToModelContent(pbBlog.Content),
	}
}

func protoBlogContentToModelContent(content []string) string {
	var res string
	for _, s := range content {
		res += s
		res += "\n"
	}
	return res
}
