package atomic_service

import (
	"atomic/atomic_model"
	"atomic/atomic_model/category"
	"atomic/atomic_store"
	"context"
)

func CommonCategoryList(ctx context.Context) ([]*category.Category, error) {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return nil, err
	}
	categories, err := category.List(ctx, db)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CategoryInsert(ctx context.Context, cg atomic_model.Category) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	return cg.Insert(ctx, db)
}

func CategoryUpdate(ctx context.Context, cg atomic_model.Category) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	return cg.Update(ctx, db)
}

func CategoryDelete(ctx context.Context, cg atomic_model.Category) error {

	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	return cg.Delete(ctx, db)
}
