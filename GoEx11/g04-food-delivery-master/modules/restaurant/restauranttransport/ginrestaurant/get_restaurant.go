package ginrestaurant

import (
	"demo/common"
	"demo/component"
	"demo/modules/restaurant/restaurantbiz"
	"demo/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
