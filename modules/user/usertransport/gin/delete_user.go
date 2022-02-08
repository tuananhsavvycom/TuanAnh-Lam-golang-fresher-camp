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

func DeleteUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewDeleteUserBiz(store)

		if err := biz.DeleteUser(c.Request.Context(), user_id); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
