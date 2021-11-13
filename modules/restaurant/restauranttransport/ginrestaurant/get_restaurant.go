package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"learn-api/common"
	"learn-api/component"
	"learn-api/modules/restaurant/restaurantbiz"
	"learn-api/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}

//type fakeCreateStore struct{}
//
//func (fakeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
//	data.Id = 10
//	return nil
//}
