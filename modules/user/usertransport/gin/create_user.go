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

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewCreateUserBiz(store)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
