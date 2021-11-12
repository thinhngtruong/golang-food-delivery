package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"learn-api/common"
	"learn-api/component"
	"learn-api/modules/restaurant/restaurantbiz"
	"learn-api/modules/restaurant/restaurantmodel"
	"learn-api/modules/restaurant/restaurantstorage"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

//type fakeCreateStore struct{}
//
//func (fakeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
//	data.Id = 10
//	return nil
//}
