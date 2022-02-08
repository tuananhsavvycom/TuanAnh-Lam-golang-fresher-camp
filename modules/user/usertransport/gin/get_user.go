package gin

import (
	"Project/common"
	"Project/component"
	"Project/modules/user/usersbiz"
	"Project/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewGetUserBiz(store)

		data, err := biz.GetUser(c.Request.Context(), user_id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
