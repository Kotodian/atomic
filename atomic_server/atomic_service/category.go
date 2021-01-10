package atomic_service

import (
	"atomic/atomic_model"
	"atomic/atomic_model/category"
	"context"
)

func (s *BlogService) CommonCategoryList(ctx context.Context) ([]*category.Category, error) {
	categories, err := category.List(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *BlogService) CategoryInsert(ctx context.Context, cg atomic_model.Category) error {

	return cg.Insert(ctx, s.DB)
}

func (s *BlogService) CategoryUpdate(ctx context.Context, cg atomic_model.Category) error {
	return cg.Update(ctx, s.DB)
}

func (s *BlogService) CategoryDelete(ctx context.Context, cg atomic_model.Category) error {

	return cg.Delete(ctx, s.DB)
}
