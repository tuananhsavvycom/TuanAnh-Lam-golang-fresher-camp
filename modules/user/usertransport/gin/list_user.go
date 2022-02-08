package gin

import (
	"Project/common"
	"Project/component"
	"Project/modules/user/usermodel"
	"Project/modules/user/usersbiz"
	"Project/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewListUserBiz(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
