package restaurantbiz

import (
	"context"
	"learn-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) SoftDeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		//return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if oldData.Status == 0 {
		//return common.ErrEntityDeleted(model.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		//return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
